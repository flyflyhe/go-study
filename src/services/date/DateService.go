package date

import (
	"time"
)

const DATE_TIME = "2006-01-02 15:04:05"
const DATE_YMD = "2006-01-02"
const ZERO_SUFFIX = " 00:00:00"
const NIGHT_SUFFIX = " 23:59:59"
const DAY_TIME = 86400

var date time.Time

func init() {
	now := time.Now()
	location, _ := time.LoadLocation("Asia/Shanghai")
	date = now.In(location)
}

func GetDateMorning(i int) string {
	tmpDate := date.AddDate(0, 0, i)
	tmpStr := tmpDate.Format(DATE_YMD)
	tmpStr = tmpStr + ZERO_SUFFIX

	return tmpStr
}

func GetDateNight(i int) string {
	tmpDate := date.AddDate(0, 0, i)
	tmpStr := tmpDate.Format(DATE_YMD)
	tmpStr = tmpStr + NIGHT_SUFFIX

	return tmpStr
}
