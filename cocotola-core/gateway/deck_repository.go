package gateway

import (
	"context"
	"errors"
	"html/template"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

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

	folderID, err := domain.NewFolderID(e.FolderID)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder id(%d): %w", e.FolderID, err)
	}

	templateID, err := domain.NewTemplateID(e.TemplateID)
	if err != nil {
		return nil, mbliberrors.Errorf("new template id(%d): %w", e.TemplateID, err)
	}

	e.Lang2 = template.HTMLEscapeString(e.Lang2)
	lang2, err := libdomain.NewLang2(e.Lang2)
	if err != nil {
		return nil, mbliberrors.Errorf("new lang2(%s): %w", e.Lang2, err)
	}

	ownerID, err := mbuserdomain.NewAppUserID(e.OwnerID)
	if err != nil {
		return nil, mbliberrors.Errorf("new app user id(%d): %w", e.OwnerID, err)
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

	deckE := DeckEntity{ //nolint:exhaustruct
		BaseModelEntity: mbusergateway.BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		SpaceID:        param.SpaceID.Int(),
		FolderID:       folderID,
		TemplateID:     param.TemplateID.Int(),
		Name:           param.Name,
		Lang2:          param.Lang2.String(),
		Description:    param.Description,
		OwnerID:        operator.AppUserID().Int(),
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

	if result := r.db.Model(
		&DeckEntity{}, //nolint:exhaustruct
	).
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

func (r *deckRepository) FindDecks(ctx context.Context, operator service.OperatorInterface, param *service.FindDecksParameter) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecks")
	defer span.End()

	spaceIDs := make([]int, 0, len(param.SpaceIDs))
	for i, id := range param.SpaceIDs {
		spaceIDs[i] = id.Int()
	}

	organizationID := uint(operator.OrganizationID().Int())

	var decksE []DeckEntity
	if result := r.db.
		Table(DeckTableName).Select(DeckTableName+".*").
		Joins("inner join "+SpaceTableName+" on "+DeckTableName+".space_id = "+SpaceTableName+".id").
		Joins("inner join "+PairOfUserAndSpaceTableName+" on "+SpaceTableName+".id = "+PairOfUserAndSpaceTableName+".space_id").
		Where(DeckTableName+".organization_id = ?", organizationID).
		Where(SpaceTableName+".organization_id = ?", organizationID).
		Where("space_id IN ?", param.SpaceIDs.IDs()).
		Where(PairOfUserAndSpaceTableName+".app_user_id = ?", operator.AppUserID().Int()).
		Find(&decksE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.FindDecks: %w", result.Error)
	}

	decks := make([]*service.Deck, 0, len(decksE))
	for i, deckE := range decksE {
		deck, err := deckE.toDeck()
		if err != nil {
			return nil, err
		}
		decks[i] = deck
	}

	return decks, nil
}

func (r *deckRepository) FindDecksByOwner(ctx context.Context, operator mbuserservice.OperatorInterface) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecksByOwner")
	defer span.End()

	var decksE []DeckEntity
	if result := r.db.
		Model(&DeckEntity{}). //nolint:exhaustruct
		Where("organization_id = ?", uint(operator.OrganizationID().Int())).
		Where("owner_id = ?", uint(operator.AppUserID().Int())).
		Find(&decksE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.FindDecksByOwner: %w", result.Error)
	}

	decks := make([]*service.Deck, 0, len(decksE))
	for _, deckE := range decksE {
		deck, err := deckE.toDeck()
		if err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}

	return decks, nil
}

func (r *deckRepository) FindDecksInPublicSpace(ctx context.Context, operator service.OperatorInterface) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecks")
	defer span.End()

	organizationID := uint(operator.OrganizationID().Int())

	var decksE []DeckEntity
	if result := r.db.WithContext(ctx).
		Table(DeckTableName).Select(DeckTableName+".*").
		Joins("inner join "+SpaceTableName+" on "+DeckTableName+".space_id = "+SpaceTableName+".id").
		Joins("inner join "+PairOfUserAndSpaceTableName+" on "+SpaceTableName+".id = "+PairOfUserAndSpaceTableName+".space_id").
		Where(DeckTableName+".organization_id = ?", organizationID).
		Where(SpaceTableName+".organization_id = ?", organizationID).
		Where(PairOfUserAndSpaceTableName+".organization_id = ?", organizationID).
		Where(PairOfUserAndSpaceTableName+".app_user_id = ?", operator.AppUserID().Int()).
		Find(&decksE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.FindDecks: %w", result.Error)
	}

	decks := make([]*service.Deck, 0, len(decksE))
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
	if result := r.db.Model(&DeckEntity{}). //nolint:exhaustruct
						Where("organization_id = ?", uint(operator.OrganizationID().Int())).Where("id = ?", deckID.Int()).
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
