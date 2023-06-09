package model

import "dream/utils/errmsg"

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// GetCate 获取全部标签
func GetCate(pageSize, pageNum int) ([]Category, int, int64) {
	var cateList []Category
	var total int64
	err = db.Find(&cateList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cateList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return cateList, errmsg.SUCCSE, total
}

// GetCateInfo 查询分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	err := db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		return cate, errmsg.ERROR
	}
	return cate, errmsg.SUCCSE
}
