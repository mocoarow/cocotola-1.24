package gateway

import (
	"context"

	"gorm.io/gorm"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbusergateway "github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type pairOfUserAndSpaceRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

type pairOfUserAndSpaceEntity struct {
	mbusergateway.JunctionModelEntity
	OrganizationID int
	AppUserID      int
	SpaceID        int
}

func (u *pairOfUserAndSpaceEntity) TableName() string {
	return "core_user_n_space"
}

func NewPairOfUserAndSpaceRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.PairOfUserAndSpaceRepository {
	return &pairOfUserAndSpaceRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *pairOfUserAndSpaceRepository) AddPairOfUserAndSpace(ctx context.Context, operator mbuserservice.AppUserInterface, appUserID *mbuserdomain.AppUserID, spaceID *domain.SpaceID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndSpaceRepository.AddPairOfUserAndSpace")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndSpaceEntity{
		JunctionModelEntity: mbusergateway.JunctionModelEntity{ //nolint:exhaustruct
			CreatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		AppUserID:      appUserID.Int(),
		SpaceID:        spaceID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return mbliberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrSpaceAlreadyExists))
	}

	return nil
}

// func (r *pairOfUserAndSpaceRepository) RemovePairOfUserAndGroup(ctx context.Context, operator service.AppUserInterface, appUserID *domain.AppUserID, userGroupID *domain.UserGroupID) error {
// 	_, span := tracer.Start(ctx, "pairOfUserAndSpaceRepository.RemovePairOfUserAndGroup")
// 	defer span.End()

// 	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.OrganizationID()}
// 	db := wrappedDB.
// 		WherePairOfUserAndGroup().
// 		Where("app_user_id = ?", appUserID.Int()).
// 		Where("user_group_id = ?", userGroupID.Int()).
// 		db
// 	result := db.Delete(&pairOfUserAndSpaceEntity{})
// 	if result.Error != nil {
// 		return mbliberrors.Errorf(". err: %w", mblibgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
// 	}
// 	if result.RowsAffected == 0 {
// 		return errors.New("ERROR")
// 	}
// 	return nil
// }

func (r *pairOfUserAndSpaceRepository) FindSpacesByUserID(_ context.Context, operator service.OperatorInterface, appUserID *mbuserdomain.AppUserID) ([]*service.Space, error) {
	spacesE := []SpaceEntity{}
	if result := r.db.Table(SpaceTableName).Select(SpaceTableName+".*").
		Where(SpaceTableName+".organization_id = ?", operator.OrganizationID().Int()).
		// Where(SpaceTableName+".removed = ?", r.dialect.BoolDefaultValue()).
		Joins("inner join "+PairOfUserAndSpaceTableName+" on "+SpaceTableName+".id = "+PairOfUserAndSpaceTableName+".space_id").
		Where(PairOfUserAndSpaceTableName+".app_user_id = ?", appUserID.Int()).
		// Order(UserGroupTableName + ".key_name").
		Find(&spacesE); result.Error != nil {
		return nil, result.Error
	}

	spaces := make([]*service.Space, len(spacesE))
	for i, e := range spacesE {
		m, err := e.toSpace()
		if err != nil {
			return nil, err
		}
		spaces[i] = m
	}

	return spaces, nil
}
