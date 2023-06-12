package model

import (
	"dream/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// GetUserById 查询单个用户
func GetUserById(id int) (User, int) {
	var user User
	err := db.Where("id = ?", id).
		Select("*").
		First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// GetUsers 搜索用户
func GetUsers(pageSize, pageNum int, userName string) ([]User, int, int64) {
	var userList []User
	var total int64
	if userName != "" {
		err := db.Select("id,username,role,created_at").
			Where("username Like ?", "%"+userName+"%").
			Limit(pageSize).Offset((pageNum - 1) * pageSize).
			Find(&userList).Error
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		db.Model(&userList).Where(
			"username LIKE ?", "%"+userName+"%").
			Count(&total)
		return userList, errmsg.SUCCSE, total
	}
	err2 := db.Select("id,username,role,created_at").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&userList).Error
	if err2 != nil {
		return nil, errmsg.ERROR, 0
	}
	db.Model(&userList).Count(&total)
	return userList, errmsg.SUCCSE, total
}

// CheckUser 检查user在不在
func CheckUser(userName string) int {
	var total int64
	db.Model(&User{}).Where("username = ?", userName).Count(&total)
	if total > 0 {
		return errmsg.ERROR_USERNAME_USED //1 "用户名已存在！
	}
	return errmsg.SUCCSE
}

// AddUser 增加用户
func AddUser(user *User) int {
	user.Password = ScryptPw(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// CheckLogin 前台登录
func CheckLogin(formDate *User) (User, int) {
	var user User
	db.Where("username = ?", formDate.Username).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formDate.Password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if err != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCSE
}

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckUpUser 更新用户检查
func CheckUpUser(id int, userName string) int {
	var user User
	db.Select("id, username").Where("username = ?", userName).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCSE
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
