package main

import (
	//"net/http"
	"subculture_ranking/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ranking", controller.GetRanking)
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
