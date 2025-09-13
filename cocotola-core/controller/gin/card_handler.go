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

	libapicard "github.com/mocoarow/cocotola-1.24/lib/api/card"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type CardQueryUsecase interface {
	FindCards(ctx context.Context, operator service.OperatorInterface) ([]*domain.CardModel, error)
}

type CardHandler struct {
	cardQueryUsecase CardQueryUsecase
	logger           *slog.Logger
}

func NewCardHandler(cardQueryUsecase CardQueryUsecase) *CardHandler {
	return &CardHandler{
		cardQueryUsecase: cardQueryUsecase,
		logger:           slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CardHandler")),
	}
}

func (h *CardHandler) FindCards(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		cards, err := h.cardQueryUsecase.FindCards(ctx, operator)
		if err != nil {
			return mbliberrors.Errorf("FindDecks: %w", err)
		}

		results := make([]libapicard.FindCardsResponseCard, len(cards))
		for i, card := range cards {
			results[i] = libapicard.FindCardsResponseCard{
				ID:         card.CardID.Int(),
				Version:    card.Version,
				TemplateID: card.TemplateID.Int(),
				Content:    card.Content,
			}
		}
		response := libapicard.FindCardsResponse{
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

func NewInitCardRouterFunc(cardQueryUsecase CardQueryUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		card := parentRouterGroup.Group("card")
		cardkHandler := NewCardHandler(cardQueryUsecase)
		for _, m := range middleware {
			card.Use(m)
		}
		card.GET("", cardkHandler.FindCards)
	}
}
