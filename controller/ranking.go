package controller

import (
	"strconv"
	"subculture_ranking/operateDb"

	"github.com/gin-gonic/gin"
)

func GetRanking(c *gin.Context) {
	var ranking operateDb.Ranking
	ranking.User_id, _ = strconv.Atoi(c.Param("user-id"))
	ranking.Category_id, _ = strconv.Atoi(c.Param("category-id"))
	// DBからselectしてJSONで返す
}
