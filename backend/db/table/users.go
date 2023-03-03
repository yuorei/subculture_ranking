package table

import "time"

// import "time"

type User struct {
	UserId          int `gorm:"primaryKey"`
	Name            string
	Age             uint
	GenreId         int `gorm:"foreignKey:CompanyRefer"`
	ProfileImageURL string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
