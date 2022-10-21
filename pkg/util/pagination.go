package util

import (
	"github.com/Dorapoketto/go-gin-example/conf"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * conf.PageSize
	}

	return result
}
