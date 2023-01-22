package db

import (
	"os"
	"subculture_ranking/db/table"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {

	var err error

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db := gorm.Openにしてしまうと再宣言となっていまいエラーを起こす
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	CreateTable()
	return db
}

func CreateTable() {
	if err := db.AutoMigrate(&table.Comment{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&table.Category{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&table.Genre{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&table.RankingList{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&table.User{}); err != nil {
		panic(err)
	}
}

func ConnectDB() *gorm.DB {
	return db
}
