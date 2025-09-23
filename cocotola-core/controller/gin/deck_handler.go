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

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
	libapicore "github.com/mocoarow/cocotola-1.24/lib/api/core"
	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/api"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin/helper"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type GuestDeckUsecase interface {
	FindDecks(ctx context.Context, operator mbuserdomain.UserInterface, param *service.FindDecksParameter) ([]*domain.Deck, error)

	// RetrieveDeckByID(ctx context.Context, operator domain.UserInterface, deckID *domain.DeckID) (*domain.Deck, error)
}
type StudentDeckUsecase interface {
	FindDecks(ctx context.Context, operator mbuserdomain.UserInterface) ([]*domain.Deck, error)

	RetrieveDeckByID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) (*domain.Deck, error)
	AddDeck(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddDeckParameter) (*domain.DeckID, error)
	UpdateDeck(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID, version int, param *service.UpdateDeckParameter) error
}

type DeckHandler struct {
	guestDeckUsecase   GuestDeckUsecase
	studentDeckUsecase StudentDeckUsecase
	rbacClient         libapi.CocotolaRBACClient
	logger             *slog.Logger
}

func NewDeckHandler(guestDeckUsecase GuestDeckUsecase, studentDeckUsecase StudentDeckUsecase, rbacClient libapi.CocotolaRBACClient) *DeckHandler {
	return &DeckHandler{
		guestDeckUsecase:   guestDeckUsecase,
		studentDeckUsecase: studentDeckUsecase,
		rbacClient:         rbacClient,
		logger:             slog.Default().With(slog.String(mbliblog.LoggerNameKey, "DeckHandler")),
	}
}

func (h *DeckHandler) FindDecks(c *gin.Context) {
	helper.HandleRoleUserFunction(c, func(ctx context.Context, operator service.RoleUserInterface) error {
		if operator.GetRole() == "guest" {
			return h.findDecksAsGuest(ctx, c, operator)
		} else if operator.GetRole() == "student" {
			return h.findDecksAsStudent(ctx, c, operator)
		}

		h.logger.WarnContext(ctx, fmt.Sprintf("invalid role: %s", operator.GetRole()))
		return mblibdomain.ErrInvalidArgument
	}, h.errorHandle)
}

func (h *DeckHandler) findDecksAsGuest(ctx context.Context, c *gin.Context, operator mbuserdomain.UserInterface) error {
	_, span := tracer.Start(ctx, "DeckHandler.findDecksAsGuest")
	defer span.End()

	var apiReq api.FindDecksRequest
	if err := c.ShouldBindQuery(&apiReq); err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return nil
	}

	spaceIDs, err := apiReq.SpaceIDs.ToSpaceIDs()
	if err != nil {
		h.logger.WarnContext(ctx, fmt.Sprintf("ToSpaceIDs: %+v", err))
		c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
		return nil
	}

	param := service.FindDecksParameter{SpaceIDs: spaceIDs}

	h.logger.InfoContext(ctx, fmt.Sprintf("findDecksAsGuest: %+v", param.SpaceIDs.IDs()))
	result, err := h.guestDeckUsecase.FindDecks(ctx, operator, &param)
	if err != nil {
		return mbliberrors.Errorf("FindDecks: %w", err)
	}

	decks := make([]api.FindDecksResponseDeck, 0, len(result))
	for _, d := range result {
		decks = append(decks, api.FindDecksResponseDeck{
			ID:          d.DeckID.Int(),
			Version:     d.Version,
			Name:        d.Name,
			Lang2:       d.Lang2.String(),
			TemplateID:  d.TemplateID.Int(),
			Description: d.Description,
		})
	}
	apiResp := api.FindDecksResponse{
		TotalCount: len(result),
		Results:    decks,
	}

	c.JSON(http.StatusOK, apiResp)
	return nil
}

func (h *DeckHandler) findDecksAsStudent(ctx context.Context, c *gin.Context, operator mbuserdomain.UserInterface) error {
	_, span := tracer.Start(ctx, "DeckHandler.findDecksAsStudent")
	defer span.End()

	result, err := h.studentDeckUsecase.FindDecks(ctx, operator)
	if err != nil {
		return mbliberrors.Errorf("FindDecks: %w", err)
	}

	c.JSON(http.StatusOK, result)
	return nil
}

func (h *DeckHandler) RetrieveDeckByID(c *gin.Context) {
	helper.HandleUserFunction(c, func(ctx context.Context, operator mbuserdomain.UserInterface) error {
		deckIDInt, err := helper.GetIntFromPath(c, "deckID")
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("GetIntFromPath: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		deckID, err := domain.NewDeckID(deckIDInt)
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("NewDeckID: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		result, err := h.studentDeckUsecase.RetrieveDeckByID(ctx, operator, deckID)
		if err != nil {
			return mbliberrors.Errorf("RetrieveDeckByID: %w", err)
		}

		c.JSON(http.StatusOK, result)
		return nil
	}, h.errorHandle)
}

func (h *DeckHandler) AddDeck(c *gin.Context) {
	helper.HandleUserFunction(c, func(ctx context.Context, operator mbuserdomain.UserInterface) error {
		var apiReq api.AddDeckRequest
		if err := c.ShouldBindJSON(&apiReq); err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("invalid parameter: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return nil
		}
		param, err := service.NewAddDeckParameter(apiReq.SpaceID.Value, apiReq.FolderID.Value, apiReq.TemplateID.Value, apiReq.Name, apiReq.Lang2.Value, apiReq.Description)
		if err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("NewAddDeckParameter: %+v", err))
			c.JSON(http.StatusBadRequest, gin.H{"message": http.StatusText(http.StatusBadRequest)})
			return nil
		}
		deckID, err := h.studentDeckUsecase.AddDeck(ctx, operator, param)
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
	helper.HandleUserFunction(c, func(ctx context.Context, operator mbuserdomain.UserInterface) error {
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

		var apiReq libapicore.UpdateDeckRequest
		if err := c.ShouldBindJSON(&apiReq); err != nil {
			h.logger.WarnContext(ctx, fmt.Sprintf("ShouldBindJSON: %+v", err))
			c.Status(http.StatusBadRequest)
			return nil
		}

		param := service.UpdateDeckParameter{
			Name:        apiReq.Name,
			Description: apiReq.Description,
		}

		if err := h.studentDeckUsecase.UpdateDeck(ctx, operator, deckID, version, &param); err != nil {
			return mbliberrors.Errorf("update deck: %w", err)
		}

		c.Status(http.StatusOK)

		return nil
	}, h.errorHandle)
}

func (h *DeckHandler) errorHandle(ctx context.Context, c *gin.Context, err error) bool {
	if errors.Is(err, mblibdomain.ErrInvalidArgument) {
		h.logger.WarnContext(ctx, fmt.Sprintf("PrivateDeckHandler: %+v", err))
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

func NewInitDeckRouterFunc(guestDeckUsecase GuestDeckUsecase, studentDeckUsecase StudentDeckUsecase, rbacClient libapi.CocotolaRBACClient) libcontroller.InitRouterGroupFunc {
	return func(parentRouterGroup gin.IRouter, middleware ...gin.HandlerFunc) {
		deck := parentRouterGroup.Group("deck")
		deckHandler := NewDeckHandler(guestDeckUsecase, studentDeckUsecase, rbacClient)
		for _, m := range middleware {
			deck.Use(m)
		}
		deck.GET("", deckHandler.FindDecks)
		deck.GET(":deckID", deckHandler.RetrieveDeckByID)
		// deck.POST(":deckID", privateDeckHandler.FindDecks)
		// deck.GET(":deckID", privateDeckHandler.FindDeckByID)
		deck.PUT(":deckID", deckHandler.UpdateDeck)
		deck.POST("", deckHandler.AddDeck)
	}
}
