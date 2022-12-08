package operateDb

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// table
type User struct {
	gorm.Model
	User_id int `gorm:"primaryKey"`
}
type Ranking struct {
	gorm.Model
	Uuid        string `json:"uuid"`
	User_id     int    `json:"user_id" gorm:"foreignKey:CompanyRefer"`
	Title       string `json:"title"`
	Image_base64 string `json:"image_base64"`
	Category_id int    `json:"category_id" gorm:"foreignKey:CompanyRefer"`
	Comment     string `json:"comment"`
	Rank        int    `json:"rank"`
}
type Category struct {
	gorm.Model
	Category_id int    `json:"Ccategory_id" gorm:"primaryKey"`
	Category    string `json:"Ccategory"`
}

func Init() error {
	return nil
}

func GetConnect() *gorm.DB {
	return db
}
