package table

import "time"

type User struct {
	UserId              int `gorm:"primaryKey"`
	Name                 string
	Age                  uint
	GenreId             int `gorm:"foreignKey:CompanyRefer"`
	ProfileImageBase64 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
