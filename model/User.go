package model

import (
	"gorm.io/gorm"
	"love_blog/utils/errmsg"
)

type User struct {
	gorm.Model
	UserId       string `gorm:"type:varchar(36);not null" json:"user_id"`
	UserName     string `gorm:"type:varchar(20);not null " json:"user_name" validate:"required,min=4,max=12" label:"用户名"`
	UserDesc     string `gorm:"type:varchar(200)" json:"user_desc"`
	UserRole     int    `gorm:"type:int;DEFAULT:2" json:"user_role" validate:"required,gte=2" label:"角色码"`
	UserEmail    string `gorm:"type:varchar(200)" json:"user_email"`
	UserImg      string `gorm:"type:varchar(200)" json:"user_img"`
	UserAvatar   string `gorm:"type:varchar(200)" json:"user_avatar"`
	IcpRecord    string `gorm:"type:varchar(200)" json:"icp_record"`
	UserPassword string `gorm:"type:varchar(500);not null" json:"user_password" validate:"required,min=6,max=120" label:"密码"`
	UserIp       string `gorm:"type:varchar(20)" json:"user_ip"`
	UserNickname string `gorm:"type:varchar(20)" json:"user_nickname"`
}

// CheckUser 检查重名
func CheckUser(userName string) int {
	var total int64
	db.Model(&User{}).Where("user_name = ?", userName).Count(&total)
	if total > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// AddUser 新增用户
func AddUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetUserInfoById 查询单个用户的详细信息
func GetUserInfoById(userId string) (User, int) {
	var userInfo User
	err := db.Where("user_id = ?", userId).First(&userInfo).Error
	if err != nil {
		return userInfo, errmsg.ERROR // 500
	}
	return userInfo, errmsg.SUCCSE
}

// CheckEditUser 更新个人信息检查
func CheckEditUser(userId string, userName string) int {
	var user User
	db.Select("user_id, user_name").Where("user_name = ?", userName).First(&user)
	code := CheckUser(userName)
	if code == errmsg.ERROR {
		return errmsg.ERROR_USERNAME_USED
	}
	if user.UserId == userId {
		return errmsg.SUCCSE
	}
	return errmsg.SUCCSE
}

// EditUser 编辑个人信息
func EditUser(userId string, user *User) int {
	maps := make(map[string]interface{})
	maps["UserName"] = user.UserName
	maps["UserRole"] = user.UserRole
	err := db.Model(&user).Where("user_id = ?", userId).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetUserList 查询用户列表 + 模糊查询
func GetUserList(userName string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	if userName == "" {
		err := db.Select("id,user_id,user_name,user_role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		if err != nil {
			return users, 0
		}
		db.Model(&users).Count(&total)
		return users, total
	}
	err2 := db.Select("id,user_id,user_name,user_role,created_at").Where("user_name LIKE ?", "%"+userName+"%").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err2 != nil {
		return users, 0
	}
	db.Model(&users).Where("user_name LIKE ?", "%"+userName+"%").Count(&total)
	return users, total
}

// DelUser 删除用户
func DelUser(userId string) int {
	err := db.Where("user_id = ?", userId).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// ChangePassword 修改密码
func ChangePassword(userId string, user *User) int {
	err := db.Model(&User{}).Where("user_id = ?", userId).Updates(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
