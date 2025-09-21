package domain

import (
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type DeckID struct {
	Value int `validate:"required,gte=1"`
}

func NewDeckID(value int) (*DeckID, error) {
	return &DeckID{
		Value: value,
	}, nil
}

func (v *DeckID) Int() int {
	return v.Value
}
func (v *DeckID) IsDeckID() bool {
	return true
}
func (v *DeckID) GetRBACObject() mbuserdomain.RBACObject {
	return mbuserdomain.NewRBACObject("deck:" + fmt.Sprint(v.Value))
}

type DeckModel struct {
	*mblibdomain.BaseModel `validate:"required"`
	DeckID                 *DeckID                      `validate:"required"`
	OrganizationID         *mbuserdomain.OrganizationID `validate:"required"`
	SpaceID                *mbuserdomain.SpaceID        `validate:"required"`
	FolderID               *FolderID                    `validate:"required"`
	Name                   string                       `validate:"required"`
	TemplateID             *TemplateID                  `validate:"required"`
	Lang2                  *mblibdomain.Lang2           `validate:"required"`
	Description            string
	OwnerID                *mbuserdomain.UserID `validate:"required"`
}

func NewDeckModel(baseModel *mblibdomain.BaseModel, deckID *DeckID, organizationID *mbuserdomain.OrganizationID, spaceID *mbuserdomain.SpaceID, folderID *FolderID, name string, templateID *TemplateID, lang2 *mblibdomain.Lang2, description string, owernID *mbuserdomain.UserID) (*DeckModel, error) {
	if baseModel == nil {
		return nil, mbliberrors.Errorf("baseModel is nil", mblibdomain.ErrInvalidArgument)
	}
	if deckID == nil {
		return nil, mbliberrors.Errorf("deckID is nil", mblibdomain.ErrInvalidArgument)
	}
	if organizationID == nil {
		return nil, mbliberrors.Errorf("organizationID is nil", mblibdomain.ErrInvalidArgument)
	}
	if spaceID == nil {
		return nil, mbliberrors.Errorf("spaceID is nil", mblibdomain.ErrInvalidArgument)
	}
	if folderID == nil {
		return nil, mbliberrors.Errorf("folderID is nil", mblibdomain.ErrInvalidArgument)
	}
	if templateID == nil {
		return nil, mbliberrors.Errorf("templateID is nil", mblibdomain.ErrInvalidArgument)
	}
	if lang2 == nil {
		return nil, mbliberrors.Errorf("lang2 is nil", mblibdomain.ErrInvalidArgument)
	}
	m := &DeckModel{
		BaseModel:      baseModel,
		DeckID:         deckID,
		OrganizationID: organizationID,
		SpaceID:        spaceID,
		FolderID:       folderID,
		Name:           name,
		TemplateID:     templateID,
		Lang2:          lang2,
		Description:    description,
		OwnerID:        owernID,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate deck model: %w", err)
	}

	return m, nil
}
