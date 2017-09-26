package traffic

import (
	"database/sql"
	"services/date"
	"services/mysql"
)

var tablename = "traffic_analysis_l3"
var db *sql.DB

type trafficAnaly struct {
	Id               int
	Vni              int
	Process          string
	Domain           string
	Ip               string
	Port             string
	Uplink_traffic   int
	Downlink_traffic int
	Start_time       string
	End_time         string
}

func init() {
	db, _ = mysql.GetMysql()
}

func GetStartId(i int) int {
	var id int
	startSql := "select id from `" + tablename + "` where start_time >= ? order by id asc limit 1"
	row := db.QueryRow(startSql, date.GetDateMorning(i))
	row.Scan(&id)

	return id
}

func GetEndId(i int) int {
	var id int
	endSql := "select id from `" + tablename + "` where start_time <= ? order by id desc limit 1"
	row := db.QueryRow(endSql, date.GetDateNight(i))
	row.Scan(&id)

	return id
}

func GetById(start int, limit int) []trafficAnaly {
	sql := "select id,vni, process,domain,ip,port,uplink_traffic,downlink_traffic,start_time,end_time from " +
		tablename + " where id > ? and port not in (53, 80) order by id asc limit ?"
	//fmt.Println(sql)
	var tas []trafficAnaly
	rows, err := db.Query(sql, start, limit)
	if err != nil {
		return tas
	}
	for rows.Next() {
		var ta trafficAnaly
		rows.Scan(&ta.Id, &ta.Vni, &ta.Process, &ta.Domain, &ta.Ip, &ta.Port, &ta.Uplink_traffic, &ta.Downlink_traffic, &ta.Start_time, &ta.End_time)
		tas = append(tas, ta)
	}

	return tas
}
