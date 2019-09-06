package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

//Datetime 格式化时间
type Datetime struct {
	time.Time
}

//MarshalJSON gin.H解析成年月日信息
func (t Datetime) MarshalJSON() ([]byte, error) {
	//格式化秒
	// seconds := t.Unix()
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

//Value 返回Time
func (t Datetime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

//Scan 验证转换方法
func (t *Datetime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Datetime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
