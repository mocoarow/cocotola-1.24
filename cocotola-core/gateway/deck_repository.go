package gateway

import (
	"context"
	"errors"
	"html/template"

	"gorm.io/gorm"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
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
	FolderID       int
	TemplateID     int
	Name           string
	Lang2          string
	Description    string
	OwnerID        int
}

func (e *DeckEntity) TableName() string {
	return "core_deck"
}

func (e *DeckEntity) toModel() (*domain.DeckModel, error) {
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

	spaceID, err := mbuserdomain.NewSpaceID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(%d): %w", e.ID, err)
	}

	folderID, err := domain.NewFolderID(e.FolderID)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder id(%d): %w", e.FolderID, err)
	}

	templateID, err := domain.NewTemplateID(e.TemplateID)
	if err != nil {
		return nil, mbliberrors.Errorf("new template id(%d): %w", e.TemplateID, err)
	}

	e.Lang2 = template.HTMLEscapeString(e.Lang2)
	lang2, err := mblibdomain.NewLang2(e.Lang2)
	if err != nil {
		return nil, mbliberrors.Errorf("new lang2(%s): %w", e.Lang2, err)
	}

	ownerID, err := mbuserdomain.NewUserID(e.OwnerID)
	if err != nil {
		return nil, mbliberrors.Errorf("new user id(%d): %w", e.OwnerID, err)
	}

	deckModel, err := domain.NewDeckModel(
		baseModel,
		deckID,
		organizationID,
		spaceID,
		folderID,
		e.Name,
		templateID,
		lang2,
		e.Description,
		ownerID,
	)
	if err != nil {
		return nil, mbliberrors.Errorf("new deck model: %w", err)
	}

	return deckModel, nil
}

func (e *DeckEntity) toDeck() (*service.Deck, error) {
	deckModel, err := e.toModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to deck model: %w", err)
	}
	deck := &service.Deck{DeckModel: deckModel}

	return deck, nil
}

type DeckEntities []DeckEntity

func (e DeckEntities) toDecks() ([]*service.Deck, error) {
	decks := make([]*service.Deck, len(e))
	for i, deckE := range e {
		deck, err := deckE.toDeck()
		if err != nil {
			return nil, mbliberrors.Errorf("to deck: %w", err)
		}
		decks[i] = deck
	}

	return decks, nil
}

type deckRepository struct {
	db *gorm.DB
}

func NewDeckRepository(db *gorm.DB) service.DeckRepository {
	return &deckRepository{
		db: db,
	}
}

func (r *deckRepository) AddDeck(ctx context.Context, operator mbuserservice.OperatorInterface, param *service.AddDeckParameter) (*domain.DeckID, error) {
	_, span := tracer.Start(ctx, "deckRepository.AddDeck")
	defer span.End()

	folderID := 0
	if param.FolderID != nil {
		folderID = param.FolderID.Int()
	}

	deckE := DeckEntity{ //nolint:exhaustruct
		BaseModelEntity: mbusergateway.BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		SpaceID:        param.SpaceID.Int(),
		FolderID:       folderID,
		TemplateID:     param.TemplateID.Int(),
		Name:           param.Name,
		Lang2:          param.Lang2.String(),
		Description:    param.Description,
		OwnerID:        operator.GetUserID().Int(),
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

func (r *deckRepository) UpdateDeck(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID, version int, param *service.UpdateDeckParameter) error {
	_, span := tracer.Start(ctx, "deckRepository.UpdateDeck")
	defer span.End()

	if result := r.db.Model(
		&DeckEntity{}, //nolint:exhaustruct
	).
		Where("organization_id = ?", uint(operator.GetOrganizationID().Int())).
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

func (r *deckRepository) FindDecks(ctx context.Context, operator mbuserservice.OperatorInterface, param *service.FindDecksParameter) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecks")
	defer span.End()

	organizationID := uint(operator.GetOrganizationID().Int())

	var decksE DeckEntities
	if result := r.db.
		Table(DeckTableName).Select(DeckTableName+".*").
		Where(DeckTableName+".organization_id = ?", organizationID).
		Where("space_id IN ?", param.SpaceIDs.IDs()).
		Find(&decksE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.FindDecks: %w", result.Error)
	}

	decks, err := decksE.toDecks()
	if err != nil {
		return nil, mbliberrors.Errorf("decksE.toDecks: %w", err)
	}

	return decks, nil
}

func (r *deckRepository) FindDecksByOwner(ctx context.Context, operator mbuserservice.OperatorInterface) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecksByOwner")
	defer span.End()

	var decksE DeckEntities
	if result := r.db.
		Model(&DeckEntity{}). //nolint:exhaustruct
		Where("organization_id = ?", uint(operator.GetOrganizationID().Int())).
		Where("owner_id = ?", uint(operator.GetUserID().Int())).
		Find(&decksE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.FindDecksByOwner: %w", result.Error)
	}

	decks, err := decksE.toDecks()
	if err != nil {
		return nil, mbliberrors.Errorf("decksE.toDecks: %w", err)
	}

	return decks, nil
}

func (r *deckRepository) RetrieveDeckByID(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID) (*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.RetrieveDeckByID")
	defer span.End()

	var deckE DeckEntity
	if result := r.db.WithContext(ctx).
		Model(&DeckEntity{}). //nolint:exhaustruct
		Where("organization_id = ?", uint(operator.GetOrganizationID().Int())).
		Where("id = ?", deckID.Int()).
		First(&deckE); result.Error != nil {
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
