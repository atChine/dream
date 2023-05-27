package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId       uint   `gorm:"type:varchar(36);not null" json:"user_id"`
	UserName     string `gorm:"type:varchar(20);not null " json:"user_name" validate:"required,min=4,max=12" label:"用户名"`
	UserDesc     string `gorm:"type:varchar(200)" json:"user_desc"`
	UserRole     int    `gorm:"type:int;DEFAULT:2" json:"user_role" validate:"required,gte=2" label:"角色码"`
	UserEmail    string `gorm:"type:varchar(200)" json:"user_email"`
	UserImg      string `gorm:"type:varchar(200)" json:"user_img"`
	UserAvatar   string `gorm:"type:varchar(200)" json:"user_avatar"`
	IcpRecord    string `gorm:"type:varchar(200)" json:"icp_record"`
	UserPassword string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	UserIp       string `gorm:"type:varchar(20)" json:"user_ip"`
	UserNickname string `gorm:"type:varchar(20)" json:"user_nickname"`
}
