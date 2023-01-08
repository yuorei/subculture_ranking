package table

import "gorm.io/gorm"

type Ranking_list struct {
	gorm.Model
	Uuid           string `gorm:"foreignKey:CompanyRefer"`
	User_id        int    ` gorm:"foreignKey:CompanyRefer"`
	Title          string
	Image_base64   string
	Category_id    int ` gorm:"foreignKey:CompanyRefer"`
	Genre_id       int ` gorm:"foreignKey:CompanyRefer"`
	Rank           int
	Poster_Comment string
}

type Category struct {
	Category_id int `gorm:"primaryKey"`
	Category    string
}

type Genre struct {
	Genre_id int    `gorm:"primaryKey"`
	Genre    string
}
