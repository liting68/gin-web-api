package model

/*
 * @Author: hiwein.lucus
 * @Date: 2019-10-12 15:04:17
 * @Last Modified by: hiwein.lucus
 * @Last Modified time: 2019-10-12 17:25:03
 */

import (
	"database/sql/driver"
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

//Datetime 格式化时间
type Datetime struct {
	time.Time
}

//UnmarshalJSON 解析格式化时间
func (d *Datetime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	sv := string(b)
	if len(sv) == 0 {
		return fmt.Errorf("%s", "can not format time empty")
	}
	if len(sv) <= 10 {
		d.Time, err = time.ParseInLocation(LayoutDate, string(sv), loc)
		if err != nil {
			d.Time, err = time.ParseInLocation(LayoutShortdate, string(sv), loc)
		}
	} else {
		d.Time, err = time.ParseInLocation(LayoutDatetime, string(sv), loc)
		if err != nil {
			d.Time, err = time.ParseInLocation(LayoutShortdateTime, string(sv), loc)
			if err != nil {
				d.Time, err = time.ParseInLocation(LayoutShortdateShortTime, string(sv), loc)
			}
		}
	}
	return err
}

//MarshalJSON gin.H解析成年月日信息
func (d Datetime) MarshalJSON() ([]byte, error) {
	//格式化秒
	if d.Unix() <= 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Format(CtLayout))), nil
}

//Value 返回Time
func (d Datetime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if d.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return d.Time, nil
}

//Scan 验证转换方法
func (d *Datetime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*d = Datetime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
