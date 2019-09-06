package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

//Timestamp 格式化时间戳
type Timestamp struct {
	time.Time
}

//MarshalJSON gin.H解析成时间戳
func (t Timestamp) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

//Value 返回Time
func (t Timestamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//Scan 验证转换方法
func (t *Timestamp) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Timestamp{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
