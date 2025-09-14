package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

func getDeckIDFromQuery(c *gin.Context) (*domain.DeckID, bool) {
	deckIDInt, err := helper.GetIntFromQuery(c, "deckID")
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil, false
	}
	deckID, err := domain.NewDeckID(deckIDInt)
	if err != nil {
		return nil, false
	}
	return deckID, true
}
