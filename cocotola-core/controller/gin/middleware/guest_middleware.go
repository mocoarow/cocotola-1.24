package middleware

import (
	"github.com/gin-gonic/gin"
)

func NewGuestMiddleware(guestOrganizationID int, guestUserID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("OrganizationID", guestOrganizationID)
		c.Set("AuthorizedUser", guestUserID)
		c.Set("Role", "guest")
	}
}
