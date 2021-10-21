package models

import (
	"blog/global"
)

type AdminRole struct {
	BaseModel
	Name string `gorm:"column:name;size:20;not null" json:"name"`
	Power string `gorm:"column:power;size:500" json:"power"`
	Desc string `gorm:"column:desc;size:255" json:"desc"`
	Status int8 `gorm:"column:status;size:4;default:0" json:"status"`
	User []AdminInfo `gorm:"foreignKey:RoleId" json:"user"`
}

func (r *AdminRole) TableName() string {
	return "admin_role"
}

func (r *AdminRole) GetPowerById(id int64) ([]int64,bool) {
	var power []int64
	result := global.DB.Model(&AdminRole{}).Where("id = ?",id).Pluck("power",&power)
	if result.RowsAffected==0{
		return nil,false
	}
	return power,true
}