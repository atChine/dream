package model

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
