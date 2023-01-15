package controller

import (
	"net/http"
	tmdbapi "subculture_ranking/TMDbAPI"
	"subculture_ranking/db"
	"subculture_ranking/db/table"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// /ranking
func GetAllRankingUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "ranking-all-users.html", gin.H{})
	var res []table.User

	db := db.ConnectDB()
	db.Find(&res)
	// todo 本当に必要なものだけを返す
	c.JSON(200, res)
}

// /ranking/:uesr-id
func GetUserRankings(c *gin.Context) {
	var res []table.RankingList

	userId := c.Param("user-id")

	db := db.ConnectDB()
	db.Find(&res, userId)

	c.JSON(200, res)
}

func PostUserRanking(c *gin.Context) {
	var res table.RankingList

	if err := c.BindJSON(&res); err != nil {
		c.JSON(400, err)
		return
	}
	uuidObj, _ := uuid.NewUUID()
	res.Uuid = uuidObj.String()
	res.CreatedAt = time.Now()
	// todo サイトからのinputで入力を引数に入れる
	data := tmdbapi.SearchTvGET(res.Title)
	res.ImageURL = "https://image.tmdb.org/t/p/w500" + data.Results[0].Posterpath
	res.Title = data.Results[0].Originalname

	db := db.ConnectDB()
	db.Create(&res)

	c.JSON(200, res)
}

// /users
func PostUserProfile(c *gin.Context) {
	var res table.User

	if err := c.BindJSON(&res); err != nil {
		c.JSON(400, err)
		return
	}

	db := db.ConnectDB()
	db.Create(&res)

	c.JSON(200, res)
}

func GetUserProfile(c *gin.Context) {
	var res table.User

	db := db.ConnectDB()
	db.Find(&res)
	c.JSON(200, res)
}
