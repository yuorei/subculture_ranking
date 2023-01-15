package db

import (
	"os"
	"subculture_ranking/db/table"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() *gorm.DB {

	var err error

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=" + os.Getenv("TIMEZONE")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
