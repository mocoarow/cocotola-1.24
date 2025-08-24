package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

func GetIntFromQuery(c *gin.Context, param string) (int, error) {
	idS := c.Query(param)
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, mbliberrors.Errorf("Atoi: %w", err)
	}

	return id, nil
}

func GetStringFromQuery(c *gin.Context, param string) string {
	return c.Query(param)
}
