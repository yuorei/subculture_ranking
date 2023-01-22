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
	r.GET("/user-register", controller.GetUserRegisterHTML)
	r.GET("/user-ranking/:user-id", controller.GetUserRankingHTML)
	r.GET("/ranking-register/:user-id",controller.GetRankingRegisterHTML)
	r.GET("/users-html/:user-id", controller.GetUserProfileHTML)

	r.GET("/ranking-user-data", controller.GetAllRankingUsers)
	r.GET("/ranking/:user-id", controller.GetUserRankings)
	r.POST("/ranking/:user-id", controller.PostUserRanking)
	r.GET("/users/:user-id", controller.GetUserProfile)
	r.POST("/users", controller.PostUserProfile)
	r.Run()
}
