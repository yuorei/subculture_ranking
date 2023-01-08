package table

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Uuid    string `gorm:"foreignKey:CompanyRefer"`
	User_id int    ` gorm:"foreignKey:CompanyRefer"`
	Comment string
}