package gateway

import (
	"context"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type DeckEntity struct {
	mbusergateway.BaseModelEntity
	ID             int
	OrganizationID int
	Name           string
	TemplateID     int
	Lang2          string
	Description    string
}

func (e *DeckEntity) TableName() string {
	return "deck"
}

func (e *DeckEntity) ToModel() (*domain.DeckModel, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, mbliberrors.Errorf("libdomain.NewBaseModel: %w", err)
	}

	organizationID, err := mbuserdomain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("domain.NewOrganizationID: %w", err)
	}

	deckModel, err := domain.NewDeckModel(
		baseModel,
		organizationID,
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
	deck := &service.Deck{
		DeckModel: deckModel,
	}
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

func (r *deckRepository) AddDeck(ctx context.Context, operator service.OperatorInterface, param *service.DeckAddParameter) (*domain.DeckID, error) {
	_, span := tracer.Start(ctx, "deckRepository.AddDeck")
	defer span.End()

	deckE := DeckEntity{
		BaseModelEntity: mbusergateway.BaseModelEntity{
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		TemplateID:     param.TemplateID,
		Name:           param.Name,
		Lang2:          param.Lang2,
		Description:    param.Description,
	}
	if result := r.db.Create(&deckE); result.Error != nil {
		return nil, mbliberrors.Errorf("deckRepository.AddDeck: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrDeckAlreadyExists))
	}

	deckID, err := domain.NewDeckID(deckE.ID)
	if err != nil {
		return nil, err
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

func (r *deckRepository) FindDecks(ctx context.Context, operator service.OperatorInterface, deckID *domain.DeckID) ([]*service.Deck, error) {
	_, span := tracer.Start(ctx, "deckRepository.FindDecks")
	defer span.End()

	var decksE []DeckEntity
	query := r.db.Model(&DeckEntity{})

	if deckID != nil {
		query = query.Where("id = ?", deckID.Int())
	}

	if result := query.Find(&decksE); result.Error != nil {
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
