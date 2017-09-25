package services

import (
	"fmt"
)

const USH = "userspeedhistory"

type uidGid struct {
	Uid int
	Gid int
}

func GetUserSpeedHistoryData() map[int]uidGid {
	data := make(map[int]uidGid)
	startDate := GetDateMorning(-2)
	endDate := GetDateNight(-1)
	//sql := "select vni,uid,gid  from `" + USH + "` where startDate > '" + startDate + "' and endDate < '" + endDate + "'"
	sql := "select vni,uid,gid  from `" + USH + "` where startDate > ? and endDate < ?"
	fmt.Println(startDate, endDate, sql)

	l3db, _ := GetL3Mysql()
	rows, err := l3db.Query(sql, startDate, endDate)
	if err != nil {
		fmt.Println(err)
	}
	if rows == nil {
		return data
	}
	for rows.Next() {
		var vni int
		var tmpUidGid uidGid
		rows.Scan(&vni, &tmpUidGid.Uid, &tmpUidGid.Gid)
		if _, ok := data[vni]; !ok {
			data[vni] = tmpUidGid
		}
	}

	return data
}
