package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIntFromQuery(c *gin.Context, param string) (int, error) {
	idS := c.Query(param)
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetStringFromQuery(c *gin.Context, param string) string {
	return c.Query(param)
}
