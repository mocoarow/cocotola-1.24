package domain

import (
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type FolderID struct {
	Value int `validate:"required,gte=1"`
}

func NewFolderID(value int) (*FolderID, error) {
	return &FolderID{
		Value: value,
	}, nil
}

func (v *FolderID) Int() int {
	return v.Value
}
func (v *FolderID) IsFolderID() bool {
	return true
}
func (v *FolderID) GetRBACObject() mbuserdomain.RBACObject {
	return mbuserdomain.NewRBACObject("folder:" + fmt.Sprint(v.Value))
}

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
	*mblibdomain.BaseModel
	DeckID         *DeckID
	OrganizationID *mbuserdomain.OrganizationID
	SpaceID        *SpaceID
	OwnerID        *mbuserdomain.AppUserID
	FolderID       *FolderID
	Name           string `validate:"required"`
	TemplateID     int    `validate:"required,gte=1"`
	Lang2          string `validate:"required"`
	Description    string
}

func NewDeckModel(baseModel *mblibdomain.BaseModel, deckID *DeckID, organizationID *mbuserdomain.OrganizationID, spaceID *SpaceID, owernID *mbuserdomain.AppUserID, folderID *FolderID, name string, templateID int, lang2 string, description string) (*DeckModel, error) {
	m := &DeckModel{
		BaseModel:      baseModel,
		OrganizationID: organizationID,
		SpaceID:        spaceID,
		OwnerID:        owernID,
		FolderID:       folderID,
		Name:           name,
		TemplateID:     templateID,
		Lang2:          lang2,
		Description:    description,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate deck model: %w", err)
	}

	return m, nil
}
