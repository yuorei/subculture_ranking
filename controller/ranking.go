package controller

import (
	"strconv"
	"subculture_ranking/operateDb"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// /ranking/:user-id/:category-id
func GetRanking(c *gin.Context) {
	var req, res operateDb.Ranking
	req.User_id, _ = strconv.Atoi(c.Param("user-id"))
	req.Category_id, _ = strconv.Atoi(c.Param("category-id"))
	// todo DBからselectしてJSONで返す
	c.JSON(200, res)
}

// /res/register
func PostRegister(c *gin.Context) {
	var req, res operateDb.Ranking
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err)
		return
	}
	// todo DBにinsert
	res.Uuid = uuid.New().String()
	res.User_id = req.User_id
	res.Category_id = req.Category_id
	res.Title = req.Title
	res.Comment = req.Comment
	res.Rank = req.Rank
	c.JSON(200, res)
}

func PutRegister(c *gin.Context) {
	var req, res operateDb.Ranking
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err)
		return
	}
	// todo DBにupdata
	res.Uuid = c.Param("uuid")
	res.User_id = req.User_id
	res.Category_id = req.Category_id
	res.Title = req.Title
	res.Comment = req.Comment
	res.Rank = req.Rank
	c.JSON(200, res)
}
