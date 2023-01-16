package main

import (
	"subculture_ranking/controller"
	"subculture_ranking/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Static("/js", "./frontend/templates/js")
	r.LoadHTMLGlob("./frontend/templates/*.html")
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	_ = db.Init()
	r.GET("/", controller.GetIndex)
	r.GET("/all-users",controller.GetAllUsersHTML)
	r.GET("/ranking", controller.GetAllRankingUsersHTML)
	r.GET("/ranking-user-data", controller.GetAllRankingUsers)
	r.GET("/ranking/:user-id", controller.GetUserRankings)
	r.POST("/ranking/", controller.PostUserRanking)
	r.GET("/users", controller.GetUserProfile)
	r.POST("/users", controller.PostUserProfile)
	r.Run()
}
