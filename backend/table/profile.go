package model

import "time"

type Profile struct {
	User_id              int ` gorm:"foreignKey:CompanyRefer"`
	Name                 string
	Age                  uint
	Genre_id             int `gorm:"foreignKey:CompanyRefer"`
	Profile_image_base64 string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
