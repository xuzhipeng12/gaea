/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 11:45 下午
 **/

package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type IdField struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}

type StartEndTime struct {
	StartTime JSONTime `gorm:"column:start_time" json:"start_time" ` // 开始时间
	EndTime   JSONTime `gorm:"column:end_time" json:"end_time"`      // 结束时间
}

type TimeField struct {
	CreatedAt JSONTime  `gorm:"column:create_time" json:"create_time"`
	UpdatedAt JSONTime  `gorm:"column:update_time" json:"update_time"`
	DeletedAt *JSONTime `gorm:"column:delete_time" json:"delete_time"`
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("无法转换 %v 的时间格式", v)
}
