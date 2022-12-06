package main

import (
	"subculture_ranking/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ranking/:user-id/:category-id", controller.GetRanking)
	r.POST("/ranking/register",controller.PostRegister)
	r.PUT("/ranking/register/:uuid",controller.PutRegister)
	r.Run()
}
