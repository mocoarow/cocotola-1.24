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

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckQueryUsecase interface {
	FindDecks(ctx context.Context, operator service.OperatorInterface) ([]*domain.DeckModel, error)

	RetrieveDeckByID(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID) (*domain.DeckModel, error)
}

type DeckCommandUsecase interface {
	AddDeck(ctx context.Context, operator service.OperatorInterface, param *service.DeckAddParameter) (*domain.DeckID, error)
	UpdateDeck(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID, version int, param *service.DeckUpdateParameter) error
}

type DeckHandler struct {
	deckQueryUsecase   DeckQueryUsecase
	deckCommandUsecase DeckCommandUsecase
	logger             *slog.Logger
}

func NewDeckHandler(deckQueryUsecase DeckQueryUsecase, deckCommandUsecase DeckCommandUsecase) *DeckHandler {
	return &DeckHandler{
		deckQueryUsecase:   deckQueryUsecase,
		deckCommandUsecase: deckCommandUsecase,
		logger:             slog.Default().With(slog.String(mbliblog.LoggerNameKey, "DeckHandler")),
	}
}

func (h *DeckHandler) FindDecks(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		// param := libapi.DeckFindParameter{
		// 	PageNo:   1,
		// 	PageSize: defaultPageSize,
		// }
		result, err := h.deckQueryUsecase.FindDecks(ctx, operator)
		if err != nil {
			return mbliberrors.Errorf("FindDecks: %w", err)
		}

		c.JSON(http.StatusOK, result)

		return nil
	}, h.errorHandle)
}

func (h *DeckHandler) RetrieveDeckByID(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		deckIDInt, err := helper.GetIntFromPath(c, "deckID")
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("GetIntFromPath. err: %+v", err))
			c.Status(http.StatusBadRequest)

			return nil
		}

		deckID, err := domain.NewDeckID(deckIDInt)
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("NewDeckID. err: %+v", err))
			c.Status(http.StatusBadRequest)

			return nil
		}

		result, err := h.deckQueryUsecase.RetrieveDeckByID(ctx, operator, deckID)
		if err != nil {
			return mbliberrors.Errorf("RetrieveDeckByID: %w", err)
		}

		c.JSON(http.StatusOK, result)

		return nil
	}, h.errorHandle)
}

func (h *DeckHandler) AddDeck(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		var apiParam libapi.DeckAddParameter
		if err := c.ShouldBindJSON(&apiParam); err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

			return nil
		}
		templateID, err := domain.NewTemplateID(apiParam.TemplateID)
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("NewTemplateID: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

			return nil
		}
		spaceID, err := domain.NewSpaceID(apiParam.SpaceID)
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("NewSpaceID: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

			return nil
		}

		param := service.DeckAddParameter{
			SpaceID:     spaceID,
			FolderID:    nil,
			TemplateID:  templateID,
			Name:        apiParam.Name,
			Lang2:       apiParam.Lang2,
			Description: apiParam.Description,
		}
		deckID, err := h.deckCommandUsecase.AddDeck(ctx, operator, &param)
		if err != nil {
			h.logger.ErrorContext(ctx, fmt.Sprintf("add deck: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})

			return nil
		}

		c.JSON(http.StatusOK, gin.H{"id": deckID.Int()})

		return nil
	}, h.errorHandle)
}

func (h *DeckHandler) UpdateDeck(c *gin.Context) {
	helper.HandleSecuredFunction(c, func(ctx context.Context, operator service.OperatorInterface) error {
		version, err := helper.GetIntFromQuery(c, "version")
		if err != nil {
			return mblibdomain.ErrInvalidArgument
		}

		deckID, err := helper.GetDeckIDFromPath(c, "deckID")
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("GetDeckIDFromPath: %+v", err))
			c.Status(http.StatusBadRequest)

			return nil
		}

		var apiParam libapi.DeckUpdateParameter
		if err := c.ShouldBindJSON(&apiParam); err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("ShouldBindJSON: %+v", err))
			c.Status(http.StatusBadRequest)

			return nil
		}

		param := service.DeckUpdateParameter{
			Name:        apiParam.Name,
			Description: apiParam.Description,
		}

		if err := h.deckCommandUsecase.UpdateDeck(ctx, operator, deckID, version, &param); err != nil {
			return mbliberrors.Errorf("update deck: %w", err)
		}

		c.Status(http.StatusOK)

		return nil
	}, h.errorHandle)
}

func (h *DeckHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	if errors.Is(err, mblibdomain.ErrInvalidArgument) {
		h.logger.WarnContext(ctx, fmt.Sprintf("PrivateDeckHandler err: %+v", err))
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

func NewInitDeckRouterFunc(deckQueryUsecase DeckQueryUsecase, deckCommandUsecase DeckCommandUsecase) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		deck := parentRouterGroup.Group("deck")
		deckHandler := NewDeckHandler(deckQueryUsecase, deckCommandUsecase)
		for _, m := range middleware {
			deck.Use(m)
		}
		deck.GET("", deckHandler.FindDecks)
		deck.GET(":deckID", deckHandler.RetrieveDeckByID)
		// deck.POST(":deckID", privateDeckHandler.FindDecks)
		// deck.GET(":deckID", privateDeckHandler.FindDeckByID)
		deck.PUT(":deckID", deckHandler.UpdateDeck)
		// deck.DELETE(":deckID", privateDeckHandler.RemoveDeck)
		deck.POST("", deckHandler.AddDeck)
	}
}
