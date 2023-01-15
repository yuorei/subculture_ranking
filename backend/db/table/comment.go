package table

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Uuid    string `gorm:"foreignKey:CompanyRefer"`
	UserId  int    ` gorm:"foreignKey:CompanyRefer"`
	Comment string
}
