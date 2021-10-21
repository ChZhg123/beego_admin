package models

type AdminPower struct {
	BaseModel
	Name string `gorm:"column:name;size:64;not null"`
	Url string `gorm:"column:url;size:128"`
	Pid int64 `gorm:"column:pid;size:20;default:0"`
	Sort int `gorm:"column:sort;size:11;default:0"`
	Icon string `gorm:"column:icon;size:60"`
	Child []AdminPower `gorm:"foreignKey:Pid;references:Id" json:"child"`
}

func (p *AdminPower) TableName() string {
	return "admin_power"
}