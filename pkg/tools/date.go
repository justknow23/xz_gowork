package tools

import (
	"fmt"
	"insurance/pkg/global"
	"math"
	"strconv"
	"strings"
	"time"
)

// GetSpecialEndDate 获取特殊字符串结束时间
// @params now string	1d,1w,1m,1y...
// @params formatString string	1d,1m,1y...
// 例如：1d 获取当天截止时间
//		1m 获取当前月截止时间
//		1y 获取当前年截止时间
func GetSpecialEndDate(Times time.Time, formatString string) (time.Time, error) {
	var resp time.Time
	sec, _ := time.ParseDuration("-1s")
	switch {
	case strings.HasSuffix(formatString, "d"):
		n, _ := strconv.Atoi(strings.TrimSuffix(formatString, "d"))
		currentYear, currentMonth, currentDay := Times.Date()
		currentLocation := Times.Location()
		firstOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
		resp = firstOfDay.AddDate(0, 0, n)
	case strings.HasSuffix(formatString, "m"):
		n, _ := strconv.Atoi(strings.TrimSuffix(formatString, "m"))
		currentYear, currentMonth, _ := Times.Date()
		currentLocation := Times.Location()
		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		resp = firstOfMonth.AddDate(0, n, 0)
	case strings.HasSuffix(formatString, "y"):
		n, _ := strconv.Atoi(strings.TrimSuffix(formatString, "y"))
		currentYear, _, _ := Times.Date()
		currentLocation := Times.Location()
		firstOfDay := time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
		resp = firstOfDay.AddDate(n, 0, 0)
	default:
		return Times, fmt.Errorf("not Supported Yet")
	}
	return resp.Add(sec), nil
}

// IsTodayFirstOfMonth 今天是否是本月的第一天
func IsTodayFirstOfMonth() bool {
	now := time.Now()
	return GetZeroTime(now).Equal(GetFirstDateOfMonth(now))
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// GetZoneTime 根据时区获取地区时间  timeZone=+6:00
func GetZoneTime(d time.Time, timeZone string) time.Time {
	if timeZone == "" || timeZone == "+08:00" {
		return d
	}
	countSplit := strings.Split(timeZone, ":")
	one := countSplit[0]
	two := countSplit[1]
	intOne, _ := strconv.ParseInt(one, 10, 64)
	intTwo, _ := strconv.ParseInt(two, 10, 64)
	resInt := intOne*3600 + intTwo*60

	totalInt := d.Unix() + resInt
	return time.Unix(totalInt, 0)
}

// TimeSub 两时间戳差几天
func TimeSub(end, begin time.Time) int {
	end = time.Date(end.Year(), end.Month(), end.Day(), end.Hour(), end.Minute(), end.Second(), 0, time.Local)
	begin = time.Date(begin.Year(), begin.Month(), begin.Day(), begin.Hour(), begin.Minute(), begin.Second(), 0, time.Local)

	floatOut := math.Ceil(end.Sub(begin).Seconds() / float64(86400))
	return int(floatOut)
}

// GetDefaultTimeStamp 1971-01-01 00:00:00时间
func GetDefaultTimeStamp() (time.Time, error) {
	return time.ParseInLocation(global.DateFmtYMDHIS, "1971-01-01 00:00:00", time.Local)
}
