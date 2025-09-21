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

type CardEntity struct {
	mbusergateway.BaseModelEntity
	ID             int
	OrganizationID int
	DeckID         int
	TemplateID     int
	Content        string
	OwnerID        int
}

func (e *CardEntity) TableName() string {
	return "core_card"
}

func (e *CardEntity) ToModel() (*domain.CardModel, error) { //nolint:dupl
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to base model: %w", err)
	}
	cardID, err := domain.NewCardID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new card id(%d): %w", e.ID, err)
	}

	organizationID, err := mbuserdomain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("new organization id(%d): %w", e.OrganizationID, err)
	}

	deckID, err := domain.NewDeckID(e.DeckID)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(%d): %w", e.DeckID, err)
	}

	templateID, err := domain.NewTemplateID(e.TemplateID)
	if err != nil {
		return nil, mbliberrors.Errorf("new template id(%d): %w", e.TemplateID, err)
	}

	ownerID, err := mbuserdomain.NewUserID(e.OwnerID)
	if err != nil {
		return nil, mbliberrors.Errorf("new user id(%d): %w", e.OwnerID, err)
	}

	cardModel, err := domain.NewCardModel(
		baseModel,
		cardID,
		organizationID,
		deckID,
		templateID,
		e.Content,
		ownerID,
	)
	if err != nil {
		return nil, mbliberrors.Errorf("new card model: %w", err)
	}

	return cardModel, nil
}

func (e *CardEntity) toCard() (*service.Card, error) {
	cardModel, err := e.ToModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to card model: %w", err)
	}
	card := &service.Card{CardModel: cardModel}

	return card, nil
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) service.CardRepository {
	return &cardRepository{
		db: db,
	}
}

func (r *cardRepository) AddCard(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddCardParameter) (*domain.CardID, error) { //nolint:dupl
	_, span := tracer.Start(ctx, "cardRepository.AddCard")
	defer span.End()

	cardE := CardEntity{ //nolint:exhaustruct
		BaseModelEntity: mbusergateway.BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		DeckID:         param.DeckID.Int(),
		TemplateID:     param.TemplateID.Int(),
		Content:        param.Content,
		OwnerID:        operator.GetUserID().Int(),
	}
	if result := r.db.Create(&cardE); result.Error != nil {
		return nil, mbliberrors.Errorf("add card entity: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrDeckAlreadyExists))
	}

	cardID, err := domain.NewCardID(cardE.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new card id(%d): %w", cardE.ID, err)
	}

	return cardID, nil
}

func (r *cardRepository) FindCardsByDeckID(ctx context.Context, operator mbuserdomain.UserInterface, deckID *domain.DeckID) ([]*service.Card, error) {
	_, span := tracer.Start(ctx, "cardRepository.FindCardsByDeckID")
	defer span.End()

	var cardsE []CardEntity
	if result := r.db.
		Model(&CardEntity{}). //nolint:exhaustruct
		Where("organization_id = ?", uint(operator.GetOrganizationID().Int())).
		Where("deck_id = ?", uint(deckID.Int())).
		Find(&cardsE); result.Error != nil {
		return nil, mbliberrors.Errorf("cardRepository.FindCardsByDeckID: %w", result.Error)
	}

	cards := make([]*service.Card, 0, len(cardsE))
	for _, cardE := range cardsE {
		card, err := cardE.toCard()
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}
