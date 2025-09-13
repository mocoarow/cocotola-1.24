package middleware

import (
	"github.com/gin-gonic/gin"
)

func NewGuestMiddleware(guestOrganizationID int, guestAppUserID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("OrganizationID", guestOrganizationID)
		c.Set("AuthorizedUser", guestAppUserID)
		c.Set("Role", "guest")
	}
}
