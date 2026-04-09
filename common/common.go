package common

import "time"

type Gender int

const (
	Female Gender = 0 //女性
	Male   Gender = 1 //男性
	Third  Gender = 2 //第三性別
)

type TimeLayout string

const (
	Date         TimeLayout = "2006-01-02"
	DateTime     TimeLayout = "2006-01-02 15:04:05"
	DateTimeZone TimeLayout = "2006-01-02 15:04:05z0700"
)

func StrToTime(str string, layout ...TimeLayout) (time.Time, error) {
	format := Date
	if len(layout) > 0 {
		format = layout[0]
	}
	return time.Parse(string(format), str)
}

func TimeToStr(t time.Time, layout ...TimeLayout) string {
	format := Date
	if len(layout) > 0 {
		format = layout[0]
	}
	return t.Format(string(format))
}
