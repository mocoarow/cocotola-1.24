package gateway

import (
	"context"
	"errors"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckEntity struct {
	mbusergateway.BaseModelEntity
	ID             int
	OrganizationID int
	SpaceID        int
	OwnerID        int
	FolderID       int
	TemplateID     int
	Name           string
	Lang2          string
	Description    string
}

func (e *DeckEntity) TableName() string {
	return "core_deck"
}

func (e *DeckEntity) ToModel() (*domain.DeckModel, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to base model: %w", err)
	}

	organizationID, err := mbuserdomain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("new organization id(%d): %w", e.OrganizationID, err)
	}

	deckID, err := domain.NewDeckID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new deck id(%d): %w", e.ID, err)
	}

	spaceID, err := domain.NewSpaceID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(%d): %w", e.ID, err)
	}

	ownerID, err := mbuserdomain.NewAppUserID(e.OwnerID)
	if err != nil {
		return nil, mbliberrors.Errorf("new app user id(%d): %w", e.OwnerID, err)
	}
	folderID, err := domain.NewFolderID(e.FolderID)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder id(%d): %w", e.FolderID, err)
	}

	deckModel, err := domain.NewDeckModel(
		baseModel,
		deckID,
		organizationID,
		spaceID,
		ownerID,
		folderID,
		e.Name,
		e.TemplateID,
		e.Lang2,
		e.Description,
	)
	if err != nil {
		return nil, mbliberrors.Errorf("new deck model: %w", err)
	}

	return deckModel, nil
}

func (e *DeckEntity) toDeck() (*service.Deck, error) {
	deckModel, err := e.ToModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to deck model: %w", err)
	}
	deck := &service.Deck{DeckModel: deckModel}

	return deck, nil
}

type deckRepository struct {
	db *gorm.DB
}

func NewDeckRepository(db *gorm.DB) service.DeckRepository {
	return &deckRepository{
		db: db,
	}
}

func (r *deckRepository) AddDeck(ctx context.Context, operator mbuserservice.OperatorInterface, param *service.DeckAddParameter) (*domain.DeckID, error) {
	_, span := tracer.Start(ctx, "deckRepository.AddDeck")
	defer span.End()

	folderID := 0
	if param.FolderID != nil {
		folderID = param.FolderID.Int()
	}

	deckE := DeckEntity{
		BaseModelEntity: mbusergateway.BaseModelEntity{
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		SpaceID:        param.SpaceID.Int(),
		OwnerID:        operator.AppUserID().Int(),
		FolderID:       folderID,
		TemplateID:     param.TemplateID.Int(),
		Name:           param.Name,
		Lang2:          param.Lang2,
		Description:    param.Description,
	}
	if result := r.db.Create(&deckE); result.Error != nil {
		return nil, mbliberrors.Errorf("add deck entity: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrDeckAlreadyExists))
	}

	deckID, err := domain.NewDeckID(deckE.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new deck id(%d): %w", deckE.ID, err)
	}

	return deckID, nil
}

func (r *deckRepository) UpdateDeck(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID, version int, param *service.DeckUpdateParameter) error {
	_, span := tracer.Start(ctx, "deckRepository.UpdateDeck")
	defer span.End()

	if result := r.db.Model(&DeckEntity{}).
		Where("organization_id = ?", uint(operator.OrganizationID().Int())).
		Where("id = ?", deckID.Int()).
		Where("version = ?", version).
		Updates(map[string]interface{}{
			"version":     gorm.Expr("version + 1"),
			"name":        param.Name,
			"description": param.Description,
		}); result.Error != nil {
		return mbliberrors.Errorf("deckRepository.UpdateDeck: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrDeckAlreadyExists))
	}

	return nil
}

func (r *deckRepository) FindDecks(ctx context.Context, operator service.OperatorInterface) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecks")
	defer span.End()

	var decksE []DeckEntity
	if result := r.db.Model(&DeckEntity{}).
		Where("organization_id = ?", uint(operator.OrganizationID().Value)).
		Find(&decksE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.FindDecks: %w", result.Error)
	}

	var decks []*service.Deck
	for _, deckE := range decksE {
		deck, err := deckE.toDeck()
		if err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}

	return decks, nil
}

func (r *deckRepository) RetrieveDeckByID(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID) (*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.RetrieveDeckByID")
	defer span.End()

	var deckE DeckEntity
	if result := r.db.Model(&DeckEntity{}).Where("organization_id = ?", uint(operator.OrganizationID().Int())).Where("id = ?", deckID.Int()).First(&deckE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrDeckNotFound
		}

		return nil, mbliberrors.Errorf("deckRepository.RetrieveDeckByID: %w", result.Error)
	}

	deck, err := deckE.toDeck()
	if err != nil {
		return nil, mbliberrors.Errorf("deckE.toDeck: %w", err)
	}

	return deck, nil
}
