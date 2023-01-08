package operateDb

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// table
type User struct {
	User_id int `gorm:"primaryKey"`
}

type Profile struct {
	User_id       int `gorm:"primaryKey"` //fkにする
	Name          string
	Age           uint
	Genre_id      int
	Profile_image int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func Init() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

func GetConnect() *gorm.DB {
	return db
}
