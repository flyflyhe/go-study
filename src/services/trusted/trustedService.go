package trusted

import (
	"encoding/json"
	"fmt"
	"services/clientGame"
	"services/date"
	"services/mysql"
	"services/redis"
	"services/slices"
	"services/traffic"
	"strconv"
	"strings"
	"sync"
)

const USH = "userspeedhistory"

type TrustedList struct {
	Pk        string
	Gid       int
	Name      string
	Area      string
	Users     int
	Process   string
	Ip        string
	Ports     string
	Flow      int
	Created   string
	Updated   string
	Real_flow int
}

type TrustedListUser struct {
	Pk         string
	Uid        int
	Start_time string
	Ports      string
	Flow       int
	Created    string
	Updated    string
}

type UidGid struct {
	Uid int
	Gid int
}

func (tl *TrustedList) addFlow(flow int) {
	tl.Flow = tl.Flow + flow
}

func (tl *TrustedList) addPorts(port string) {
	tl.Ports = tl.Ports + ";" + port
}

func (tl *TrustedList) mergePorts(ports string) {
	s1 := strings.Split(tl.Ports, ";")
	s2 := strings.Split(ports, ";")
	s3 := slices.Merge(s1, s2)
	tl.Ports = strings.Join(s3, ";")
}
func (tl *TrustedList) addUsers(user int) {
	tl.Users = tl.Users + user
}

func (tlu *TrustedListUser) addFlow(flow int) {
	tlu.Flow = tlu.Flow + flow
}

func (tlu *TrustedListUser) addPorts(port string) {
	tlu.Ports = tlu.Ports + ";" + port
}

func (tlu *TrustedListUser) mergePorts(ports string) {
	s1 := strings.Split(tlu.Ports, ";")
	s2 := strings.Split(ports, ";")
	s3 := slices.Merge(s1, s2)
	tlu.Ports = strings.Join(s3, ";")
}

func GetUserSpeedHistoryData() map[int]UidGid {
	data := make(map[int]UidGid)
	startDate := date.GetDateMorning(-2)
	endDate := date.GetDateNight(-1)
	sql := "select vni,uid,gid  from `" + USH + "` where startDate > ? and endDate < ?"
	fmt.Println(startDate, endDate, sql)

	l3db, _ := mysql.GetL3Mysql()
	rows, err := l3db.Query(sql, startDate, endDate)
	if err != nil {
		fmt.Println(err)
	}
	if rows == nil {
		return data
	}
	for rows.Next() {
		var vni int
		var tmpUidGid UidGid
		rows.Scan(&vni, &tmpUidGid.Uid, &tmpUidGid.Gid)
		if _, ok := data[vni]; !ok {
			data[vni] = tmpUidGid
		}
	}

	return data
}

func Run() {
	startDate := date.GetDate()
	var wg sync.WaitGroup
	limit := 20000
	ushData := GetUserSpeedHistoryData()
	gameData := clientGame.GetGame()
	startId := traffic.GetStartId(-1)
	endId := traffic.GetEndId(-1)
	times := CoumputeTimes(startId, endId, limit)
	todey := date.GetDateMorning(-1)
	wg.Add(times)

	fmt.Println(startId, ";", endId, ";", times)

	trustedlist := make(chan map[string]*TrustedList, times)
	trustedlistuser := make(chan map[string]*TrustedListUser, times)

	counter := 0
	for ; startId < endId; startId = startId + limit {
		counter++
		if (endId - startId) < limit {
			limit = endId - startId
		}
		tasData := traffic.GetById(startId, limit)
		//fmt.Println(startId)
		println(counter)
		go func() {
			defer wg.Done()

			trustedlistTmp := make(map[string]*TrustedList)
			trustedlistuserTmp := make(map[string]*TrustedListUser)

			for _, v := range tasData {
				if v.Process == "" || v.Ip == "" {
					continue
				}
				if _, ok := ushData[v.Vni]; !ok {
					continue
				}
				if _, ok := gameData[ushData[v.Vni].Gid]; !ok {
					continue
				}
				var tl TrustedList
				var tlu TrustedListUser

				uid := ushData[v.Vni].Uid
				gid := ushData[v.Vni].Gid
				ipc := string([]byte(v.Ip)[0:strings.LastIndex(v.Ip, ".")])
				pk := strings.TrimSpace(strconv.Itoa(gid)) + "_" + strings.ToLower(v.Process) + "_" + ipc
				uip := strings.TrimSpace(strconv.Itoa(uid)) + "_" + v.Ip
				sum := v.Uplink_traffic + v.Downlink_traffic

				if sum == 0 {
					continue
				}

				if _, ok := trustedlistTmp[pk]; ok {
					trustedlistTmp[pk].addFlow(sum)
					trustedlistTmp[pk].addUsers(1)
					if strings.Index(trustedlistTmp[pk].Ports, v.Port) == -1 {
						trustedlistTmp[pk].addPorts(v.Port)
					}
				} else {
					tl.Pk = pk
					tl.Flow = sum
					tl.Gid = gid
					tl.Area = gameData[gid].Area
					tl.Name = gameData[gid].Name
					tl.Ip = v.Ip
					tl.Process = v.Process
					tl.Ports = v.Port
					tl.Users = 1
					tl.Real_flow = 0
					tl.Created = todey

					trustedlistTmp[pk] = &tl
				}
				//fmt.Println(trustedlistTmp[pk])

				if _, ok := trustedlistuserTmp[uip]; ok {
					if trustedlistuserTmp[uip].Pk == pk {
						if strings.Index(trustedlistuserTmp[uip].Ports, v.Port) == -1 {
							trustedlistuserTmp[uip].addPorts(v.Port)
						}
						trustedlistuserTmp[uip].addFlow(sum)
					}
				} else {
					tlu.Flow = sum
					tlu.Pk = pk
					tlu.Ports = v.Port
					tlu.Start_time = v.Start_time
					tlu.Uid = uid
					tlu.Created = todey
					tlu.Updated = todey

					trustedlistuserTmp[uip] = &tlu
				}
			}

			//println(startId)
			trustedlist <- trustedlistTmp
			trustedlistuser <- trustedlistuserTmp
		}()
	}

	wg.Wait()
	fmt.Println("wait结束")
	close(trustedlist)
	close(trustedlistuser)

	wg.Add(2)
	trustedlistFinal := make(map[string]*TrustedList)
	trustedlistuserFinal := make(map[string]*TrustedListUser)

	fmt.Println("数据计算开始")
	go func() {
		defer wg.Done()
		for tlmapTmp := range trustedlist {
			for k, v := range tlmapTmp {
				if _, ok := trustedlistFinal[k]; ok {
					trustedlistFinal[k].addFlow(v.Flow)
					trustedlistFinal[k].addUsers(v.Users)
					trustedlistFinal[k].mergePorts(v.Ports)
				} else {
					trustedlistFinal[k] = v
				}
			}
		}

		fmt.Println("6库长度")
		fmt.Println(len(trustedlistFinal))

		rc := redis.GetRedis6()
		for pk, tl := range trustedlistFinal {
			if tl.Flow < 100000 {
				continue
			}
			rc.HSet(pk, "pk", tl.Pk)
			rc.HSet(pk, "flow", tl.Flow)
			rc.HSet(pk, "gid", tl.Gid)
			rc.HSet(pk, "area", tl.Area)
			rc.HSet(pk, "name", tl.Name)
			rc.HSet(pk, "ip", tl.Ip)
			rc.HSet(pk, "process", tl.Process)
			rc.HSet(pk, "ports", tl.Ports)
			rc.HSet(pk, "users", tl.Users)
			rc.HSet(pk, "created", tl.Created)
		}
		fmt.Println("redis6计算完成")
	}()

	go func() {
		defer wg.Done()
		for tlumapTmp := range trustedlistuser {
			for k, v := range tlumapTmp {
				if _, ok := trustedlistuserFinal[k]; ok {
					trustedlistuserFinal[k].addFlow(v.Flow)
					trustedlistuserFinal[k].mergePorts(v.Ports)
				} else {
					trustedlistuserFinal[k] = v
				}
				//fmt.Println(v)
			}
		}

		fmt.Println("7库长度")
		fmt.Println(len(trustedlistuserFinal))

		rc2 := redis.GetRedis7()
		rc2.FlushDB()
		for uip, tlu := range trustedlistuserFinal {
			pk := tlu.Pk
			jsonData, _ := json.Marshal(tlu)
			rc2.HSet(pk, uip, jsonData)
			// rc2.HSet(pk, "flow", tlu.Flow)
			// rc2.HSet(pk, "pk", tlu.Pk)
			// rc2.HSet(pk, "ports", tlu.Ports)
			// rc2.HSet(pk, "start_time", tlu.Start_time)
			// rc2.HSet(pk, "uid", tlu.Uid)
		}
		fmt.Println("redis7计算完成")
	}()
	wg.Wait()
	date.Init()
	endDate := date.GetDate()
	fmt.Println("开始时间" + startDate + "结束时间" + endDate)
	fmt.Println("结束")
}

func CoumputeTimes(start int, end int, limit int) int {
	var i = 0
	for ; start < end; start = start + limit {
		if (end - start) < limit {
			limit = end - start
		}
		i = i + 1
	}
	return i
}
