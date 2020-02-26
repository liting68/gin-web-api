package utils

import (
	"fmt"
	"time"
)

//格式化时间
const (
	LayoutDatetime           = "2006-01-02 15:04:05"
	LayoutShortdateTime      = "2006-1-2 15:04:05"
	LayoutShortdateShortTime = "2006-1-2 15:4:5"
	LayoutDateMin            = "2006-01-02 15:04"
	LayoutDate               = "2006-01-02"
	LayoutShortdate          = "2006-1-2"
)

//StringToTime 格式化时间
func StringToTime(timestr string) (t time.Time, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	sv := string(timestr)
	if len(sv) == 0 {
		return t, fmt.Errorf("%s", "can not format time empty")
	}
	if len(sv) <= 10 {
		t, err = time.ParseInLocation(LayoutDate, string(sv), loc)
		if err != nil {
			t, err = time.ParseInLocation(LayoutShortdate, string(sv), loc)
		}
	} else {
		t, err = time.ParseInLocation(LayoutDatetime, string(sv), loc)
		if err != nil {
			t, err = time.ParseInLocation(LayoutShortdateTime, string(sv), loc)
			if err != nil {
				t, err = time.ParseInLocation(LayoutShortdateShortTime, string(sv), loc)
			}
		}
	}
	return t, err
}

//FormatDatetime 时间格式化字符2006-01-02 15:04:05
func FormatDatetime(t time.Time) string {
	return t.Format(LayoutDatetime)
}

//FormatDate 时间格式化字符2006-01-02
func FormatDate(t time.Time) string {
	return t.Format(LayoutDate)
}
