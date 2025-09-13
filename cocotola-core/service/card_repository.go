package service

import (
	"context"
	"errors"

	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrCardAlreadyExists = errors.New("card already exists")
var ErrCardNotFound = errors.New("card not found")

type AddCardParameter struct {
	DeckID     *domain.DeckID
	TemplateID *domain.TemplateID
	Content    string
}

type CardRepository interface {
	AddCard(ctx context.Context, operator mbuserservice.OperatorInterface, param *AddCardParameter) (*domain.CardID, error)

	FindCardsByDeckID(ctx context.Context, operator mbuserservice.OperatorInterface, deckID *domain.DeckID) ([]*Card, error)
}
