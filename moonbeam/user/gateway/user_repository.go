package gateway

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	libgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	"github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	"github.com/mocoarow/cocotola-1.24/moonbeam/user/service"
)

type userEntity struct {
	BaseModelEntity
	ID                   int
	OrganizationID       int
	LoginID              string
	Username             string
	HashedPassword       string
	Provider             string
	ProviderID           string
	ProviderAccessToken  string
	ProviderRefreshToken string
	Deleted              bool
}

func (e *userEntity) TableName() string {
	return UserTableName
}

func (e *userEntity) toUser(userGroups []*domain.UserGroupModel) (*domain.User, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, liberrors.Errorf("e.toModel. err: %w", err)
	}

	userID, err := domain.NewUserID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewUser. err: %w", err)
	}

	organizationID, err := domain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOrganizationID. err: %w", err)
	}

	userModel, err := domain.NewUser(baseModel, userID, organizationID, e.LoginID, e.Username, userGroups)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewUser. err: %w", err)
	}

	return userModel, nil
}

func (e *userEntity) toOwnerModel(userGroups []*domain.UserGroupModel) (*domain.OwnerModel, error) {
	userModel, err := e.toUser(userGroups)
	if err != nil {
		return nil, liberrors.Errorf("e.toUser. err: %w", err)
	}

	ownerModel, err := domain.NewOwnerModel(userModel)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOwnerModel. err: %w", err)
	}

	return ownerModel, nil
}

func (e *userEntity) toSystemOwnerModel(_ context.Context, _ service.RepositoryFactory, userGroup []*domain.UserGroupModel) (*domain.SystemOwner, error) {
	if e.LoginID != service.SystemOwnerLoginID {
		return nil, liberrors.Errorf("invalid system owner. loginID: %s", e.LoginID)
	}

	ownerModel, err := e.toOwnerModel(userGroup)
	if err != nil {
		return nil, liberrors.Errorf("e.toOwnerModel(). err: %w", err)
	}

	systemOwnerModel, err := domain.NewSystemOwnerModel(ownerModel)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewSystemOwnerModel. err: %w", err)
	}

	return systemOwnerModel, nil
}

type userRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

var _ service.UserRepository = (*userRepository)(nil)

func NewUserRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.UserRepository {
	return &userRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *userRepository) FindSystemOwnerByOrganizationID(ctx context.Context, _ domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.SystemOwner, error) {
	_, span := tracer.Start(ctx, "userRepository.FindSystemOwnerByOrganizationID")
	defer span.End()

	var user userEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereUser().Where(UserTableName+".login_id = ?", service.SystemOwnerLoginID).db
	if result := db.First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, liberrors.Errorf("system owner not found. organization ID: %d, err: %w", organizationID, service.ErrSystemOwnerNotFound)
		}

		return nil, result.Error
	}

	return user.toSystemOwnerModel(ctx, r.rf, nil)
}

func (r *userRepository) FindSystemOwnerByOrganizationName(ctx context.Context, _ domain.SystemAdminInterface, organizationName string, options ...service.Option) (*domain.SystemOwner, error) {
	_, span := tracer.Start(ctx, "userRepository.FindSystemOwnerByOrganizationName")
	defer span.End()

	var userE userEntity
	if result := r.db.Table(OrganizationTableName).Select(UserTableName+".*").
		Where(OrganizationTableName+".name = ? and "+UserTableName+".deleted = ?", organizationName, r.dialect.BoolDefaultValue()).
		Where("login_id = ?", service.SystemOwnerLoginID).
		Joins("inner join " + UserTableName + " on " + OrganizationTableName + ".id = " + UserTableName + ".organization_id").
		First(&userE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, liberrors.Errorf("system owner not found. organization name: %s, err: %w", organizationName, service.ErrSystemOwnerNotFound)
		}

		return nil, result.Error
	}

	user, err := userE.toUser(nil)
	if err != nil {
		return nil, err
	}

	userGroups := []*domain.UserGroupModel{}
	for _, option := range options {
		if option == service.IncludeGroups {
			pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, r.dialect, r.db, r.rf)
			userGroupsTmp, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, user, user.GetUserID())
			if err != nil {
				return nil, liberrors.Errorf("FindUserGroupsByUserID: %w", err)
			}

			userGroups = userGroupsTmp
		}
	}

	return userE.toSystemOwnerModel(ctx, r.rf, userGroups)
}

func (r *userRepository) FindUserByID(ctx context.Context, operator domain.UserInterface, id *domain.UserID, options ...service.Option) (*domain.User, error) {
	_, span := tracer.Start(ctx, "userRepository.FindUserByID")
	defer span.End()

	return r.findUserByID(ctx, operator.GetOrganizationID(), id, options...)
}

func (r *userRepository) findUserByID(ctx context.Context, organizationID *domain.OrganizationID, id *domain.UserID, options ...service.Option) (*domain.User, error) {
	_, span := tracer.Start(ctx, "userRepository.findUserByID")
	defer span.End()

	var userE userEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereUser().Where(UserTableName+".id = ?", id.Int()).db
	if result := db.First(&userE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrUserNotFound
		}

		return nil, result.Error
	}

	user, err := userE.toUser(nil)
	if err != nil {
		return nil, liberrors.Errorf("toUser: %w", err)
	}

	userGroups := []*domain.UserGroupModel{}

	for _, option := range options {
		if option == service.IncludeGroups {
			pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, r.dialect, r.db, r.rf)
			userGroupsTmp, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, user, user.GetUserID())
			if err != nil {
				return nil, liberrors.Errorf("FindUserGroupsByUserID: %w", err)
			}

			userGroups = userGroupsTmp
		}
	}

	return userE.toUser(userGroups)
}

func (r *userRepository) FindUserByLoginID(ctx context.Context, operator domain.UserInterface, loginID string) (*domain.User, error) {
	_, span := tracer.Start(ctx, "userRepository.FindUserByLoginID")
	defer span.End()

	return r.findUserByLoginID(ctx, operator.GetOrganizationID(), loginID)
}

func (r *userRepository) findUserByLoginID(ctx context.Context, organizationID *domain.OrganizationID, loginID string) (*domain.User, error) {
	_, span := tracer.Start(ctx, "userRepository.findUserByLoginID")
	defer span.End()

	userEntity, err := r.findUserEntityByLoginID(ctx, organizationID, loginID)
	if err != nil {
		return nil, err
	}

	return userEntity.toUser(nil)
}

func (r *userRepository) findUserEntityByLoginID(ctx context.Context, organizationID *domain.OrganizationID, loginID string) (*userEntity, error) {
	_, span := tracer.Start(ctx, "userRepository.findUserEntityByLoginID")
	defer span.End()

	var user userEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereUser().Where(UserTableName+".login_id = ?", loginID).db
	if result := db.First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrUserNotFound
		}

		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) FindOwnerByLoginID(ctx context.Context, operator domain.SystemOwnerInterface, loginID string) (*domain.OwnerModel, error) {
	_, span := tracer.Start(ctx, "userRepository.FindOwnerByLoginID")
	defer span.End()

	var user userEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.GetOrganizationID()}
	db := wrappedDB.Table(UserTableName).Select(UserTableName+".*").
		WherePairOfUserAndGroup().
		WhereUserGroup().
		WhereUser().
		Where(UserTableName+".login_id = ?", loginID).
		Where(UserGroupTableName+".key_name = ? ", service.OwnerGroupKey).
		Joins("inner join " + PairOfUserAndGroupTableName + " on " + UserTableName + ".id = " + PairOfUserAndGroupTableName + ".user_id").
		Joins("inner join " + UserGroupTableName + " on " + PairOfUserAndGroupTableName + ".user_group_id = " + UserGroupTableName + ".id").
		db

	if result := db.First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrUserNotFound
		}

		return nil, result.Error
	}

	return user.toOwnerModel(nil)
}

func (r *userRepository) addUser(ctx context.Context, userEntity *userEntity) (*domain.UserID, error) {
	_, span := tracer.Start(ctx, "userRepository.addUser")
	defer span.End()

	if result := r.db.Create(userEntity); result.Error != nil {
		return nil, liberrors.Errorf("db.Create. err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	userID, err := domain.NewUserID(userEntity.ID)
	if err != nil {
		return nil, liberrors.Errorf("NewUserID: %w", err)
	}

	return userID, nil
}

func (r *userRepository) AddUser(ctx context.Context, operator domain.UserInterface, param *service.AddUserParameter) (*domain.UserID, error) {
	_, span := tracer.Start(ctx, "userRepository.AddUser")
	defer span.End()

	hashedPassword := ""
	if len(param.Password) != 0 {
		hashedPasswordTmp, err := libgateway.HashPassword(param.Password)
		if err != nil {
			return nil, liberrors.Errorf("libgateway.HashPassword. err: %w", err)
		}

		hashedPassword = hashedPasswordTmp
	}

	userEntity := userEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: operator.GetOrganizationID().Int(),
		LoginID:        param.LoginID,
		Username:       param.Username,
		HashedPassword: hashedPassword,
	}

	userID, err := r.addUser(ctx, &userEntity)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func (r *userRepository) AddSystemOwner(ctx context.Context, operator domain.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserID, error) {
	_, span := tracer.Start(ctx, "userRepository.AddSystemOwner")
	defer span.End()

	userEntity := userEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: organizationID.Int(),
		LoginID:        service.SystemOwnerLoginID,
		Username:       "SystemOwner",
	}

	userID, err := r.addUser(ctx, &userEntity)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

func (r *userRepository) VerifyPassword(ctx context.Context, operator domain.SystemOwnerInterface, loginID, password string) (bool, error) {
	organizationID := operator.GetOrganizationID()
	userEntity, err := r.findUserEntityByLoginID(ctx, organizationID, loginID)
	if err != nil {
		return false, err
	}

	return ComparePasswords(userEntity.HashedPassword, password), nil
}

func ComparePasswords(hashedPassword string, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}
