package gateway

import (
	"context"

	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type pairOfUserAndSpaceRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
}

type pairOfUserAndSpaceEntity struct {
	JunctionModelEntity
	OrganizationID int
	UserID         int
	SpaceID        int
}

func (u *pairOfUserAndSpaceEntity) TableName() string {
	return PairOfUserAndSpaceTableName
}

func NewPairOfUserAndSpaceRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB) service.PairOfUserAndSpaceRepository {
	return &pairOfUserAndSpaceRepository{
		dialect: dialect,
		db:      db,
	}
}

func (r *pairOfUserAndSpaceRepository) AddPairOfUserAndSpace(ctx context.Context, operator service.UserInterface, userID *domain.UserID, spaceID *domain.SpaceID) error {
	_, span := tracer.Start(ctx, "pairOfUserAndSpaceRepository.AddPairOfUserAndSpace")
	defer span.End()

	pairOfUserAndGroup := pairOfUserAndSpaceEntity{
		JunctionModelEntity: JunctionModelEntity{ //nolint:exhaustruct
			CreatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		UserID:         userID.Int(),
		SpaceID:        spaceID.Int(),
	}
	if result := r.db.Create(&pairOfUserAndGroup); result.Error != nil {
		return liberrors.Errorf(". err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	return nil
}

func (r *pairOfUserAndSpaceRepository) FindMySpaces(ctx context.Context, operator service.UserInterface) ([]*service.Space, error) {
	spacesE := []spaceEntity{}
	if result := r.db.WithContext(ctx).Table(SpaceTableName).Select(SpaceTableName+".*").
		Where(SpaceTableName+".organization_id = ?", operator.GetOrganizationID().Int()).
		Where(SpaceTableName+".deleted = ?", r.dialect.BoolDefaultValue()).
		Where(UserTableName+".organization_id = ?", operator.GetOrganizationID().Int()).
		Where(UserTableName+".id = ?", operator.GetUserID().Int()).
		Where(UserTableName+".deleted = ?", r.dialect.BoolDefaultValue()).
		Joins("inner join " + PairOfUserAndSpaceTableName + " on " + SpaceTableName + ".id = " + PairOfUserAndSpaceTableName + ".space_id").
		Joins("inner join " + UserTableName + " on " + PairOfUserAndSpaceTableName + ".user_id = " + UserTableName + ".id").
		Order(SpaceTableName + ".key_name").
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
