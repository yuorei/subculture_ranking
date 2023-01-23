package table

import "gorm.io/gorm"

type RankingList struct {
	gorm.Model
	Uuid          string `gorm:"primaryKey" gorm:"foreignKey:CompanyRefer"`
	UserId        int    ` gorm:"foreignKey:CompanyRefer"`
	Title         string `gorm:"unique;not null"`
	ImageURL      string `gorm:"unique"`
	CategoryId    int    ` gorm:"foreignKey:CompanyRefer"`
	GenreId       int    ` gorm:"foreignKey:CompanyRefer"`
	Rank          int    `gorm:"unique;not null"`
	PosterComment string
}

type Category struct {
	CategoryId int `gorm:"primaryKey"`
	Category   string
}

type Genre struct {
	GenreId int `gorm:"primaryKey"`
	Genre   string
}
