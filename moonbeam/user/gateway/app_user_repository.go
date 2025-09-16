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

type appUserRepository struct {
	dialect libgateway.DialectRDBMS
	db      *gorm.DB
	rf      service.RepositoryFactory
}

type appUserEntity struct {
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

func (e *appUserEntity) TableName() string {
	return UserTableName
}

// func (e *appUserEntity) toUser(ctx context.Context, rf service.RepositoryFactory, userGroups []domain.UserGroupModel) (*service.User, error) {
// 	appUserModel, err := e.toUserModel(userGroups)
// 	if err != nil {
// 		return nil, err
// 	}
// 	appUser, err := service.NewUser(ctx, rf, appUserModel)
// 	if err != nil {
// 		return nil, err

//		}
//		return appUser, nil
//	}
func (e *appUserEntity) toUserModel(userGroups []*domain.UserGroupModel) (*domain.UserModel, error) {
	baseModel, err := e.ToBaseModel()
	if err != nil {
		return nil, liberrors.Errorf("e.toModel. err: %w", err)
	}

	appUserID, err := domain.NewUserID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewUserModel. err: %w", err)
	}

	organizationID, err := domain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOrganizationID. err: %w", err)
	}

	appUserModel, err := domain.NewUserModel(baseModel, appUserID, organizationID, e.LoginID, e.Username, userGroups)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewUserModel. err: %w", err)
	}

	return appUserModel, nil
}

func (e *appUserEntity) toOwnerModel(userGroups []*domain.UserGroupModel) (*domain.OwnerModel, error) {
	appUserModel, err := e.toUserModel(userGroups)
	if err != nil {
		return nil, liberrors.Errorf("e.toUserModel. err: %w", err)
	}

	ownerModel, err := domain.NewOwnerModel(appUserModel)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOwnerModel. err: %w", err)
	}

	return ownerModel, nil
}

func (e *appUserEntity) toSystemOwner(ctx context.Context, rf service.RepositoryFactory, userGroup []*domain.UserGroupModel) (*service.SystemOwner, error) {
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

	systemOwner, err := service.NewSystemOwner(ctx, rf, systemOwnerModel)
	if err != nil {
		return nil, liberrors.Errorf("service.NewSystemOwner. err: %w", err)
	}

	return systemOwner, nil
}

func (e *appUserEntity) toOwner(rf service.RepositoryFactory, userGroup []*domain.UserGroupModel) (*service.Owner, error) {
	ownerModel, err := e.toOwnerModel(userGroup)
	if err != nil {
		return nil, liberrors.Errorf("e.toOwnerModel(). err: %w", err)
	}

	return service.NewOwner(rf, ownerModel), nil
}

func (e *appUserEntity) toUser(ctx context.Context, rf service.RepositoryFactory, userGroups []*domain.UserGroupModel) (*service.User, error) {
	appUserModel, err := e.toUserModel(userGroups)
	if err != nil {
		return nil, liberrors.Errorf("e.toUserModel(). err: %w", err)
	}

	appUser, err := service.NewUser(ctx, rf, appUserModel)
	if err != nil {
		return nil, liberrors.Errorf("service.NewUser. err: %w", err)
	}

	return appUser, nil
}

func NewUserRepository(_ context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.UserRepository {
	return &appUserRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *appUserRepository) FindSystemOwnerByOrganizationID(ctx context.Context, _ service.SystemAdminInterface, organizationID *domain.OrganizationID) (*service.SystemOwner, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindSystemOwnerByOrganizationID")
	defer span.End()

	var appUser appUserEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereUser().Where(UserTableName+".login_id = ?", service.SystemOwnerLoginID).db
	if result := db.First(&appUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, liberrors.Errorf("system owner not found. organization ID: %d, err: %w", organizationID, service.ErrSystemOwnerNotFound)
		}

		return nil, result.Error
	}

	return appUser.toSystemOwner(ctx, r.rf, nil)
}

func (r *appUserRepository) FindSystemOwnerByOrganizationName(ctx context.Context, _ service.SystemAdminInterface, organizationName string, options ...service.Option) (*service.SystemOwner, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindSystemOwnerByOrganizationName")
	defer span.End()

	var appUserE appUserEntity
	if result := r.db.Table(OrganizationTableName).Select(UserTableName+".*").
		Where(OrganizationTableName+".name = ? and "+UserTableName+".deleted = ?", organizationName, r.dialect.BoolDefaultValue()).
		Where("login_id = ?", service.SystemOwnerLoginID).
		Joins("inner join " + UserTableName + " on " + OrganizationTableName + ".id = " + UserTableName + ".organization_id").
		First(&appUserE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, liberrors.Errorf("system owner not found. organization name: %s, err: %w", organizationName, service.ErrSystemOwnerNotFound)
		}

		return nil, result.Error
	}

	appUser, err := appUserE.toUser(ctx, r.rf, nil)
	if err != nil {
		return nil, err
	}

	userGroups := []*domain.UserGroupModel{}
	for _, option := range options {
		if option == service.IncludeGroups {
			pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, r.dialect, r.db, r.rf)
			userGroupsTmp, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, appUser, appUser.GetUserID())
			if err != nil {
				return nil, liberrors.Errorf("FindUserGroupsByUserID: %w", err)
			}

			userGroups = userGroupsTmp
		}
	}

	return appUserE.toSystemOwner(ctx, r.rf, userGroups)
}

func (r *appUserRepository) FindUserByID(ctx context.Context, operator service.UserInterface, id *domain.UserID, options ...service.Option) (*service.User, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindUserByID")
	defer span.End()

	return r.findUserByID(ctx, operator.GetOrganizationID(), id, options...)
}

func (r *appUserRepository) findUserByID(ctx context.Context, organizationID *domain.OrganizationID, id *domain.UserID, options ...service.Option) (*service.User, error) {
	_, span := tracer.Start(ctx, "appUserRepository.findUserByID")
	defer span.End()

	var appUserE appUserEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereUser().Where(UserTableName+".id = ?", id.Int()).db
	if result := db.First(&appUserE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrUserNotFound
		}

		return nil, result.Error
	}

	appUser, err := appUserE.toUser(ctx, r.rf, nil)
	if err != nil {
		return nil, liberrors.Errorf("toUser: %w", err)
	}

	userGroups := []*domain.UserGroupModel{}

	for _, option := range options {
		if option == service.IncludeGroups {
			pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, r.dialect, r.db, r.rf)
			userGroupsTmp, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, appUser, appUser.GetUserID())
			if err != nil {
				return nil, liberrors.Errorf("FindUserGroupsByUserID: %w", err)
			}

			userGroups = userGroupsTmp
		}
	}

	return appUserE.toUser(ctx, r.rf, userGroups)
}

func (r *appUserRepository) FindUserByLoginID(ctx context.Context, operator service.UserInterface, loginID string) (*service.User, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindUserByLoginID")
	defer span.End()

	return r.findUserByLoginID(ctx, operator.GetOrganizationID(), loginID)
}

func (r *appUserRepository) findUserByLoginID(ctx context.Context, organizationID *domain.OrganizationID, loginID string) (*service.User, error) {
	_, span := tracer.Start(ctx, "appUserRepository.findUserByLoginID")
	defer span.End()

	appUserEntity, err := r.findUserEntityByLoginID(ctx, organizationID, loginID)
	if err != nil {
		return nil, err
	}

	return appUserEntity.toUser(ctx, r.rf, nil)
}

func (r *appUserRepository) findUserEntityByLoginID(ctx context.Context, organizationID *domain.OrganizationID, loginID string) (*appUserEntity, error) {
	_, span := tracer.Start(ctx, "appUserRepository.findUserEntityByLoginID")
	defer span.End()

	var appUser appUserEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereUser().Where(UserTableName+".login_id = ?", loginID).db
	if result := db.First(&appUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrUserNotFound
		}

		return nil, result.Error
	}

	return &appUser, nil
}

func (r *appUserRepository) FindOwnerByLoginID(ctx context.Context, operator service.SystemOwnerInterface, loginID string) (*service.Owner, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindOwnerByLoginID")
	defer span.End()

	var appUser appUserEntity
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.GetOrganizationID()}
	db := wrappedDB.Table(UserTableName).Select(UserTableName+".*").
		WherePairOfUserAndGroup().
		WhereUserGroup().
		WhereUser().
		Where(UserTableName+".login_id = ?", loginID).
		Where(UserGroupTableName+".key_name = ? ", service.OwnerGroupKey).
		Joins("inner join " + PairOfUserAndGroupTableName + " on " + UserTableName + ".id = " + PairOfUserAndGroupTableName + ".app_user_id").
		Joins("inner join " + UserGroupTableName + " on " + PairOfUserAndGroupTableName + ".user_group_id = " + UserGroupTableName + ".id").
		db

	if result := db.First(&appUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrUserNotFound
		}

		return nil, result.Error
	}

	return appUser.toOwner(r.rf, nil)
}

func (r *appUserRepository) addUser(_ context.Context, appUserEntity *appUserEntity) (*domain.UserID, error) {
	if result := r.db.Create(appUserEntity); result.Error != nil {
		return nil, liberrors.Errorf("db.Create. err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrUserAlreadyExists))
	}

	appUserID, err := domain.NewUserID(appUserEntity.ID)
	if err != nil {
		return nil, liberrors.Errorf("NewUserID: %w", err)
	}

	return appUserID, nil
}

func (r *appUserRepository) AddUser(ctx context.Context, operator service.OwnerModelInterface, param *service.AddUserParameter) (*domain.UserID, error) {
	_, span := tracer.Start(ctx, "appUserRepository.AddUser")
	defer span.End()

	hashedPassword := ""
	if len(param.Password) != 0 {
		hashedPasswordTmp, err := libgateway.HashPassword(param.Password)
		if err != nil {
			return nil, liberrors.Errorf("libgateway.HashPassword. err: %w", err)
		}

		hashedPassword = hashedPasswordTmp
	}

	appUserEntity := appUserEntity{ //nolint:exhaustruct
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

	appUserID, err := r.addUser(ctx, &appUserEntity)
	if err != nil {
		return nil, err
	}

	return appUserID, nil
}

func (r *appUserRepository) AddSystemOwner(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.UserID, error) {
	_, span := tracer.Start(ctx, "appUserRepository.AddSystemOwner")
	defer span.End()

	appUserEntity := appUserEntity{ //nolint:exhaustruct
		BaseModelEntity: BaseModelEntity{ //nolint:exhaustruct
			Version:   1,
			CreatedBy: operator.GetUserID().Int(),
			UpdatedBy: operator.GetUserID().Int(),
		},
		OrganizationID: organizationID.Int(),
		LoginID:        service.SystemOwnerLoginID,
		Username:       "SystemOwner",
	}

	appUserID, err := r.addUser(ctx, &appUserEntity)
	if err != nil {
		return nil, err
	}

	return appUserID, nil
}

func (r *appUserRepository) VerifyPassword(ctx context.Context, operator service.SystemOwnerInterface, loginID, password string) (bool, error) {
	organizationID := operator.GetOrganizationID()
	appUserEntity, err := r.findUserEntityByLoginID(ctx, organizationID, loginID)
	if err != nil {
		return false, err
	}

	return ComparePasswords(appUserEntity.HashedPassword, password), nil
}

func ComparePasswords(hashedPassword string, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}

	return true
}
