package models

import (
	"blog/global"
)

type AdminInfo struct {
	BaseModel
	UserName string `gorm:"column:username;unique;size:20;not null" json:"username"`
	NickName string `gorm:"column:nickname;size:20" json:"nickname"`
	Phone string `gorm:"column:phone;size:20" json:"phone"`
	PassWord string `gorm:"column:password;size:128" json:"password"`
	RoleId int64 `gorm:"column:role_id;size:11;default:0" json:"role_id"`
	Status int8 `gorm:"column:status;size:4;default:0" json:"status"`
	LoginTime MyTime `gorm:"column:login_time;type:datetime" json:"login_time"`
	LoginIp string `gorm:"column:login_ip;size(50)" json:"login_ip"`
	Role AdminRole `json:"role" gorm:"foreignKey:role_id"`
}

func (admin *AdminInfo) TableName() string {
	return "admin_info"
}

func (admin *AdminInfo) GetUserByUserName(username string) *AdminInfo {
	result := global.DB.Model(&AdminInfo{}).Preload("Role").Where("username = ?",username).First(&admin)
	if result.RowsAffected == 0{
		return nil
	}
	return admin
}






