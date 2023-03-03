package db

import (
	"fmt"
	"os"
	"subculture_ranking/db/table"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {

	var err error
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	TZ := os.Getenv("TZ")

	dsn := "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=5432 sslmode=disable TimeZone=" + TZ

	//// dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DB") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db := gorm.Openにしてしまうと再宣言となっていまいエラーを起こす
	//dsn := "host=db user=postgres password=postgres dbname=subculture_ranking port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB接続失敗 " + err.Error())
	} else {
		fmt.Println("DB接続はOK")
	}

// 	CreateTable()
	
// }

// func CreateTable() {
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
	fmt.Println("テーブルOK")
	return db
}

func ConnectDB() *gorm.DB {
	return db
}
