package table

import "time"

type User struct {
	User_id              int `gorm:"primaryKey"`
	Name                 string
	Age                  uint
	Genre_id             int `gorm:"foreignKey:CompanyRefer"`
	Profile_image_base64 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
