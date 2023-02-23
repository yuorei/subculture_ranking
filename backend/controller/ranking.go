package controller

import (
	"net/http"
	"strconv"
	tmdbapi "subculture_ranking/TMDbAPI"
	"subculture_ranking/db"
	"subculture_ranking/db/table"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// /
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// /all-users
func GetAllUsersHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "user-profile.html", gin.H{})
}

// /ranking
func GetAllRankingUsersHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "ranking-all-users.html", gin.H{})
}

// /user-register
func GetUserRegisterHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "user-register.html", gin.H{})
}

// /ranking-register/:user-id
func GetRankingRegisterHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "ranking-register.html", gin.H{})
}

// /user-ranking/:user-id
func GetUserRankingHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "ranking-user.html", gin.H{})
}

// /users-html/:user-id
func GetUserProfileHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "user-profile.html", gin.H{})
}

// /ranking-data
func GetAllRankingUsers(c *gin.Context) {
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
	intUserId, _ := strconv.Atoi(userId)
	db.Order("rank ASC").Where("user_id = ?", intUserId).Find(&res)

	c.JSON(200, res)
}

// /ranking/:user-id
func PostUserRanking(c *gin.Context) {
	userId := c.Param("user-id")
	var res table.RankingList
	if err := c.BindJSON(&res); err != nil {
		c.JSON(400, err)
		return
	}
	res.UserId, _ = strconv.Atoi(userId)
	uuidObj, _ := uuid.NewUUID()
	res.Uuid = uuidObj.String()
	res.CreatedAt = time.Now()
	data, err := tmdbapi.SearchTvGET(res.Title)
	if err != nil {
		c.JSON(400, res)
		return
	}
	res.ImageURL = "https://image.tmdb.org/t/p/w200" + data.Results[0].Posterpath
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

// /users/:user-id
func GetUserProfile(c *gin.Context) {
	var res table.User
	userId := c.Param("user-id")
	db := db.ConnectDB()
	intUserId, _ := strconv.Atoi(userId)
	db.Where("user_id = ?", intUserId).Find(&res)
	c.JSON(200, res)
}
