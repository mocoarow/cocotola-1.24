package gateway

import (
	"context"
	"errors"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type FolderEntity struct {
	mbusergateway.BaseModelEntity
	ID             int
	OrganizationID int
	SpaceID        int
	ParentID       int
	Name           string
	OwnerID        int
}

func (e *FolderEntity) TableName() string {
	return "core_folder"
}

func (e *FolderEntity) toFolder() (*domain.Folder, error) { //nolint:dupl
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, mbliberrors.Errorf("to base model: %w", err)
	}

	organizationID, err := mbuserdomain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, mbliberrors.Errorf("new organization id(%d): %w", e.OrganizationID, err)
	}

	folderID, err := domain.NewFolderID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder id(%d): %w", e.ID, err)
	}

	spaceID, err := mbuserdomain.NewSpaceID(e.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new space id(%d): %w", e.ID, err)
	}

	parentID, err := domain.NewFolderID(e.ParentID)
	if err != nil {
		return nil, mbliberrors.Errorf("new parent id(%d): %w", e.ParentID, err)
	}

	ownerID, err := mbuserdomain.NewUserID(e.OwnerID)
	if err != nil {
		return nil, mbliberrors.Errorf("new user id(%d): %w", e.OwnerID, err)
	}

	folder, err := domain.NewFolder(
		baseModel,
		folderID,
		organizationID,
		spaceID,
		parentID,
		e.Name,
		ownerID,
	)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder model: %w", err)
	}

	return folder, nil
}

type folderRepository struct {
	db *gorm.DB
}

func NewFolderRepository(db *gorm.DB) service.FolderRepository {
	return &folderRepository{
		db: db,
	}
}

func (r *folderRepository) RetrieveRooFolderBySpaceID(ctx context.Context, operator mbuserdomain.UserInterface, spaceID *mbuserdomain.SpaceID) (*domain.Folder, error) {
	_, span := tracer.Start(ctx, "folderRepository.FindRooFolderBySpaceID")
	defer span.End()

	var folderE FolderEntity
	if result := r.db.WithContext(ctx).
		Where("organization_id = ?", operator.GetOrganizationID().Int()).
		Where("space_id = ? ", spaceID.Int()).
		Where("parent_id = 0").
		First(&folderE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrFolderNotFound
		}
		return nil, mbliberrors.Errorf("find folder entity by id(%d): %w", spaceID.Int())
	}

	folder, err := folderE.toFolder()
	if err != nil {
		return nil, mbliberrors.Errorf("to folder: %w", err)
	}

	return folder, nil
}

func (r *folderRepository) AddFolder(ctx context.Context, operator mbuserdomain.UserInterface, param *service.AddFolderParameter) (*domain.FolderID, error) { //nolint:dupl
	_, span := tracer.Start(ctx, "folderRepository.AddFolder")
	defer span.End()

	folderE := FolderEntity{ //nolint:exhaustruct
		BaseModelEntity: mbusergateway.BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		SpaceID:        param.SpaceID.Int(),
		ParentID:       param.FolderID.Int(),
		Name:           param.Name,
		OwnerID:        operator.GetUserID().Int(),
	}
	if result := r.db.Create(&folderE); result.Error != nil {
		return nil, mbliberrors.Errorf("add folder entity: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrFolderAlreadyExists))
	}

	folderID, err := domain.NewFolderID(folderE.ID)
	if err != nil {
		return nil, mbliberrors.Errorf("new folder id(%d): %w", folderE.ID, err)
	}

	return folderID, nil
}
