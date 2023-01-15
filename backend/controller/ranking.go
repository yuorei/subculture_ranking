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


func GetUserRanking(c *gin.Context) {
	var res table.RankingList

	db := db.ConnectDB()
	db.Find(&res)
	c.JSON(200, res)
}

// /users
func PostUserProfile(c *gin.Context) {
	var res table.User
	db := db.ConnectDB()
	if err := c.BindJSON(&res);err!=nil{
		c.JSON(400, err)
		return
	}
	db.Create(&res)
	c.JSON(200,res)
}

func GetUserProfile(c *gin.Context){
	var res table.User

	db := db.ConnectDB()
	db.Find(&res)
	c.JSON(200, res)
}

