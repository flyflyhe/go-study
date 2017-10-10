package date

import (
	"fmt"
	"os"
	"time"
)

const DATE_TIME = "2006-01-02 15:04:05"
const DATE_YMD = "2006-01-02"
const ZERO_SUFFIX = " 00:00:00"
const NIGHT_SUFFIX = " 23:59:59"
const DAY_TIME = 86400

var dateIns time.Time

func init() {
	TrueInit()
}

func Reinit() {
	TrueInit()
}

func TrueInit() {
	now := time.Now()
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	dateIns = now.In(location)
}

func GetDateMorning(i int) string {
	tmpDate := dateIns.AddDate(0, 0, i)
	tmpStr := tmpDate.Format(DATE_YMD)
	tmpStr = tmpStr + ZERO_SUFFIX

	return tmpStr
}

func GetDateNight(i int) string {
	tmpDate := dateIns.AddDate(0, 0, i)
	tmpStr := tmpDate.Format(DATE_YMD)
	tmpStr = tmpStr + NIGHT_SUFFIX

	return tmpStr
}

func GetDate() string {
	tmpStr := dateIns.Format(DATE_TIME)

	return tmpStr
}
