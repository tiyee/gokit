package helps

import (
	"strings"
	"time"
)

const OneDaySec = int64(86400)

var offset int = 8 * 3600

func init() {
	_, offset = time.Now().Zone()
}

func FormatTime(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format(time.DateTime)
}

func FormatDate(ts int64) string {
	t := time.Unix(ts, 0)
	return t.Format(time.DateOnly)
}

func StrToTime(str string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(time.DateTime, str, loc)
	if err != nil {
		return 0
	}
	return theTime.Unix()
}

func Str2Time(str string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(time.DateTime, str, loc)
	return theTime
}
func Duration(s, e int64) string {
	var start string
	var end string
	if s == 0 {
		start = ""
	} else {
		start = time.Unix(int64(s), 0).Format(time.DateTime)
	}
	if e == 0 {
		end = ""
	} else {
		end = time.Unix(int64(e), 0).Format(time.DateTime)
	}
	return start + " - " + end
}
func Duration2(start, end int64) string {
	return strings.Join([]string{
		time.Unix(start, 0).Format(time.DateTime),
		" - ",
		time.Unix(end, 0).Format("15:05"),
	}, "")
}
func TodayFirstSecond() int64 {
	return FirstSecond(time.Now())
}
func FirstSecond(t time.Time) int64 {
	//return (t.Unix()/(24*3600))*24*3600 - 8*3600
	_, offset := t.Zone()
	return getFirstSecond(t.Unix(), offset)
}
func FirstSecond2(s int64) int64 {
	return getFirstSecond(s, offset)
}
func DurationSecondsOfMonth(s int64) (int64, int64) {
	t := time.Unix(s, 0)
	return GetFirstDateOfMonth(t).Unix(), GetLastDateOfMonth(t).Unix()
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的23:59:59。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	d = time.Unix(FirstSecond(d), 0)
	return d.AddDate(0, 1, -1)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
func getFirstSecond(ts int64, offset int) int64 {
	off := int64(offset)
	return ((ts+off)/OneDaySec)*OneDaySec - off
}
