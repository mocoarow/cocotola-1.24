package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	libapicore "github.com/mocoarow/cocotola-1.24/lib/api/core"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type CardUsecase interface {
	FindCardsByDeckID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*domain.Card, error)
}

type CardHandler struct {
	cardUsecase CardUsecase
	logger      *slog.Logger
}

func NewCardHandler(cardQueryUsecase CardUsecase) *CardHandler {
	return &CardHandler{
		cardUsecase: cardQueryUsecase,
		logger:      slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CardHandler")),
	}
}

func (h *CardHandler) FindCards(c *gin.Context) {
	helper.HandleUserFunction(c, func(ctx context.Context, operator mbuserdomain.UserInterface) error {
		deckID, ok := getDeckIDFromQuery(c)
		if !ok {
			return nil
		}

		cards, err := h.cardUsecase.FindCardsByDeckID(ctx, operator, deckID)
		if err != nil {
			return mbliberrors.Errorf("FindDecks: %w", err)
		}

		results := make([]libapicore.FindCardsResponseCard, len(cards))
		for i, card := range cards {
			results[i] = libapicore.FindCardsResponseCard{
				ID:         card.CardID.Int(),
				Version:    card.Version,
				TemplateID: card.TemplateID.Int(),
				Content:    card.Content,
			}
		}
		response := libapicore.FindCardsResponse{
			TotalCount: len(results),
			Results:    results,
		}

		c.JSON(http.StatusOK, response)
		return nil
	}, h.errorHandle)
}

func (h *CardHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	if errors.Is(err, mblibdomain.ErrInvalidArgument) {
		h.logger.WarnContext(ctx, fmt.Sprintf("CardHandler err: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

		return true
	}
	if errors.Is(err, service.ErrDeckNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})

		return true
	}
	h.logger.ErrorContext(ctx, fmt.Sprintf("DeckHandler. error: %+v", err))

	return false
}

func NewInitCardRouterFunc(cardUsecase CardUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		card := parentRouterGroup.Group("card")
		cardkHandler := NewCardHandler(cardUsecase)
		for _, m := range middleware {
			card.Use(m)
		}
		card.GET("", cardkHandler.FindCards)
	}
}
