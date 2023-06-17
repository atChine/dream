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

// AddCategory 增加分类标签
func AddCategory(category *Category) int {
	err := db.Create(category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckCategory 查询分类重复
func CheckCategory(name string) int {
	var cate Category
	db.Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// EditCate 修改标签
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
