package model

import (
	"gorm.io/gorm"
	"love_blog/utils/errmsg"
)

type Category struct {
	gorm.Model
	CategoryId int    `gorm:"type:int;not null" json:"category_id"`
	Name       string `gorm:"type:varchar(20);not null" json:"name"`
}

// AddCate 新增分类标签
func AddCate(cate *Category) int {
	err := db.Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckCate 查看分类标签
func CheckCate(cateName string) int {
	var cate Category
	if cateName == "" {
		return errmsg.ERROR
	}
	db.Select("id").Where("name = ?", cateName).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

// DelCateById 根据id删除分类标签
func DelCateById(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// EditCateById 根据id编辑分类名字
func EditCateById(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetCateInfoById 通过id查询单个分类详细信息
func GetCateInfoById(id int) (Category, int) {
	var cate Category
	err := db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCSE
}

// GetCateList 查询分类列表
func GetCateList(pageSize int, pageNum int) ([]Category, int64, int) {
	var cate []Category
	var total int64
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageNum).Find(&cate).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	db.Model(&cate).Count(&total)
	return cate, total, errmsg.SUCCSE
}
