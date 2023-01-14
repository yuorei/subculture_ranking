package controller

import (

	"subculture_ranking/db"
	"subculture_ranking/db/table"


	"github.com/gin-gonic/gin"
)

// /ranking
func GetAllRankings(c *gin.Context) {
	var res []table.User

	db := db.ConnectDB()
	db.Find(&res)
	// todo 本当に必要なものだけを返す
	c.JSON(200, res)
}
