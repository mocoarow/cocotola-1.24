package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

func GetIntFromPath(c *gin.Context, param string) (int, error) {
	idS := c.Param(param)
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, mbliberrors.Errorf("Atoi. param: %s, value: %s, err: %w", param, idS, err)
	}

	return id, nil
}

func GetStringFromPath(c *gin.Context, param string) string {
	return c.Param(param)
}
