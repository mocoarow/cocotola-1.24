package domain

import (
	"fmt"

	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type FolderID struct {
	Value int `validate:"gte=0"`
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

var EmptyFolderID = &FolderID{Value: 0}

type FolderModel struct {
	*mblibdomain.BaseModel `validate:"required"`
	FolderID               *FolderID                    `validate:"required"`
	OrganizationID         *mbuserdomain.OrganizationID `validate:"required"`
	SpaceID                *mbuserdomain.SpaceID        `validate:"required"`
	ParentID               *FolderID                    `validate:"required"`
	Name                   string                       `validate:"required"`
	OwnerID                *mbuserdomain.UserID      `validate:"required"`
}

func NewFolderModel(baseModel *mblibdomain.BaseModel, folderID *FolderID, organizationID *mbuserdomain.OrganizationID, spaceID *mbuserdomain.SpaceID, parentID *FolderID, name string, owernID *mbuserdomain.UserID) (*FolderModel, error) {
	if baseModel == nil {
		return nil, mbliberrors.Errorf("baseModel is nil", mblibdomain.ErrInvalidArgument)
	}
	if folderID == nil {
		return nil, mbliberrors.Errorf("folderID is nil", mblibdomain.ErrInvalidArgument)
	}
	if organizationID == nil {
		return nil, mbliberrors.Errorf("organizationID is nil", mblibdomain.ErrInvalidArgument)
	}
	if spaceID == nil {
		return nil, mbliberrors.Errorf("spaceID is nil", mblibdomain.ErrInvalidArgument)
	}
	if parentID == nil {
		return nil, mbliberrors.Errorf("parentID is nil", mblibdomain.ErrInvalidArgument)
	}
	m := &FolderModel{
		BaseModel:      baseModel,
		FolderID:       folderID,
		OrganizationID: organizationID,
		SpaceID:        spaceID,
		ParentID:       folderID,
		Name:           name,
		OwnerID:        owernID,
	}

	if err := mblibdomain.Validator.Struct(m); err != nil {
		return nil, mbliberrors.Errorf("validate folder model: %w", err)
	}

	return m, nil
}
