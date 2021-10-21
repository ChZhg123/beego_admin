package models

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type BaseModel struct {
	Id        int64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreateTime MyTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime MyTime `gorm:"column:update_time" json:"update_time"`
}

func (v BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("create_time",time.Now())
	tx.Statement.SetColumn("update_time",time.Now())
	return nil
}

func (v BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("update_time",time.Now())
	return
}

func Status(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?",0)
}

func Paginate(page string,perPage string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(perPage)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

type MyTime struct {
	time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
