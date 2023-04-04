package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type SystemTime struct {
	time.Time
}

const (
	// 时间序列化格式
	JSON_TIME_FORMAT    = "2006-01-02 15:04:05"
	// 日期序列化格式
	JSON_DATE_FORMAT    = "2006-01-02"
	// 时区
	LOCATION            = "Asia/Shanghai"
	// GMT时差
	TIME_DELTA          = 8 * time.Hour
)

//
// MarshalJSON
//  @Description: json序列化方法，用于响应输出
//  @receiver a 注意使用值类型而非指针类型
//  @return []byte
//  @return error
//
func (a SystemTime) MarshalJSON() ([]byte, error) {
	realTime := a.Add(TIME_DELTA)
	var stamp = fmt.Sprintf("\"%s\"", realTime.Format(JSON_TIME_FORMAT))
	return []byte(stamp), nil

}

//
// UnmarshalJSON
//  @Description: json反序列化方法，用于接收输入
//  @receiver a
//  @param data
//  @return error
//
func (a *SystemTime) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	// 调整时区
	location, _ := time.LoadLocation(LOCATION)
	res, err := time.ParseInLocation(JSON_TIME_FORMAT, value, location)
	if err != nil {
		res, _ = time.ParseInLocation(JSON_DATE_FORMAT, value, location)
	}
	a.Time = res
	return nil
}

//
// Value
//  @Description: 写入数据库时调用的方法
//  @receiver a 注意使用值类型而非指针类型
//  @return driver.Value
//  @return error
//
func (a SystemTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if a.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return a.Time, nil
}

//
// Scan
//  @Description: 从数据库中读取时调用的方法
//  @receiver a
//  @param data
//  @return error
//
func (a *SystemTime) Scan(data interface{}) error {
	value, ok := data.(time.Time)
	if ok {
		*a = SystemTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", data)
}

