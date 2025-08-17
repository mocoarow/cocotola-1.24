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
	Removed              bool
}

func (e *appUserEntity) TableName() string {
	return AppUserTableName
}

// func (e *appUserEntity) toAppUser(ctx context.Context, rf service.RepositoryFactory, userGroups []domain.UserGroupModel) (*service.AppUser, error) {
// 	appUserModel, err := e.toAppUserModel(userGroups)
// 	if err != nil {
// 		return nil, err
// 	}
// 	appUser, err := service.NewAppUser(ctx, rf, appUserModel)
// 	if err != nil {
// 		return nil, err

//		}
//		return appUser, nil
//	}
func (e *appUserEntity) toAppUserModel(userGroups []*domain.UserGroupModel) (*domain.AppUserModel, error) {
	baseModel, err := e.toBaseModel()
	if err != nil {
		return nil, liberrors.Errorf("e.toModel. err: %w", err)
	}

	appUserID, err := domain.NewAppUserID(e.ID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewAppUserModel. err: %w", err)
	}

	organizationID, err := domain.NewOrganizationID(e.OrganizationID)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewOrganizationID. err: %w", err)
	}

	appUserModel, err := domain.NewAppUserModel(baseModel, appUserID, organizationID, e.LoginID, e.Username, userGroups)
	if err != nil {
		return nil, liberrors.Errorf("domain.NewAppUserModel. err: %w", err)
	}

	return appUserModel, nil
}

func (e *appUserEntity) toOwnerModel(userGroups []*domain.UserGroupModel) (*domain.OwnerModel, error) {
	appUserModel, err := e.toAppUserModel(userGroups)
	if err != nil {
		return nil, liberrors.Errorf("e.toAppUserModel. err: %w", err)
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

func (e *appUserEntity) toAppUser(ctx context.Context, rf service.RepositoryFactory, userGroups []*domain.UserGroupModel) (*service.AppUser, error) {
	appUserModel, err := e.toAppUserModel(userGroups)
	if err != nil {
		return nil, liberrors.Errorf("e.toAppUserModel(). err: %w", err)
	}

	appUser, err := service.NewAppUser(ctx, rf, appUserModel)
	if err != nil {
		return nil, liberrors.Errorf("service.NewAppUser. err: %w", err)
	}

	return appUser, nil
}

func NewAppUserRepository(ctx context.Context, dialect libgateway.DialectRDBMS, db *gorm.DB, rf service.RepositoryFactory) service.AppUserRepository {
	return &appUserRepository{
		dialect: dialect,
		db:      db,
		rf:      rf,
	}
}

func (r *appUserRepository) FindSystemOwnerByOrganizationID(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID) (*service.SystemOwner, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindSystemOwnerByOrganizationID")
	defer span.End()

	appUser := appUserEntity{}
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereAppUser().Where(AppUserTableName+".login_id = ?", service.SystemOwnerLoginID).db
	if result := db.First(&appUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, liberrors.Errorf("system owner not found. organization ID: %d, err: %w", organizationID, service.ErrSystemOwnerNotFound)
		}
		return nil, result.Error
	}

	return appUser.toSystemOwner(ctx, r.rf, nil)
}

func (r *appUserRepository) FindSystemOwnerByOrganizationName(ctx context.Context, operator service.SystemAdminInterface, organizationName string, options ...service.Option) (*service.SystemOwner, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindSystemOwnerByOrganizationName")
	defer span.End()

	appUserE := appUserEntity{}
	if result := r.db.Table(OrganizationTableName).Select(AppUserTableName+".*").
		Where(OrganizationTableName+".name = ? and "+AppUserTableName+".removed = ?", organizationName, r.dialect.BoolDefaultValue()).
		Where("login_id = ?", service.SystemOwnerLoginID).
		Joins("inner join " + AppUserTableName + " on " + OrganizationTableName + ".id = " + AppUserTableName + ".organization_id").
		First(&appUserE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, liberrors.Errorf("system owner not found. organization name: %s, err: %w", organizationName, service.ErrSystemOwnerNotFound)
		}

		return nil, result.Error
	}

	appUser, err := appUserE.toAppUser(ctx, r.rf, nil)
	if err != nil {
		return nil, err
	}

	userGroups := []*domain.UserGroupModel{}
	for _, option := range options {
		if option == service.IncludeGroups {
			pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, r.dialect, r.db, r.rf)
			userGroupsTmp, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, appUser, appUser.AppUserID())
			if err != nil {
				return nil, err
			}

			userGroups = userGroupsTmp
		}
	}

	return appUserE.toSystemOwner(ctx, r.rf, userGroups)
}

func (r *appUserRepository) FindAppUserByID(ctx context.Context, operator service.AppUserInterface, id *domain.AppUserID, options ...service.Option) (*service.AppUser, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindAppUserByID")
	defer span.End()

	return r.findAppUserByID(ctx, operator.OrganizationID(), id, options...)
}

func (r *appUserRepository) findAppUserByID(ctx context.Context, organizationID *domain.OrganizationID, id *domain.AppUserID, options ...service.Option) (*service.AppUser, error) {
	_, span := tracer.Start(ctx, "appUserRepository.findAppUserByID")
	defer span.End()

	appUserE := appUserEntity{}
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereAppUser().Where(AppUserTableName+".id = ?", id.Int()).db
	if result := db.First(&appUserE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrAppUserNotFound
		}

		return nil, result.Error
	}

	appUser, err := appUserE.toAppUser(ctx, r.rf, nil)
	if err != nil {
		return nil, err
	}

	userGroups := []*domain.UserGroupModel{}

	for _, option := range options {
		if option == service.IncludeGroups {
			pairOfUserAndGroupRepo := NewPairOfUserAndGroupRepository(ctx, r.dialect, r.db, r.rf)
			userGroupsTmp, err := pairOfUserAndGroupRepo.FindUserGroupsByUserID(ctx, appUser, appUser.AppUserID())
			if err != nil {
				return nil, err
			}

			userGroups = userGroupsTmp
		}
	}

	return appUserE.toAppUser(ctx, r.rf, userGroups)
}

func (r *appUserRepository) FindAppUserByLoginID(ctx context.Context, operator service.AppUserInterface, loginID string) (*service.AppUser, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindAppUserByLoginID")
	defer span.End()

	return r.findAppUserByLoginID(ctx, operator.OrganizationID(), loginID)
}

func (r *appUserRepository) findAppUserByLoginID(ctx context.Context, organizationID *domain.OrganizationID, loginID string) (*service.AppUser, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindAppUserByLoginID")
	defer span.End()

	appUserEntity, err := r.findAppUserEntityByLoginID(ctx, organizationID, loginID)
	if err != nil {
		return nil, err
	}

	return appUserEntity.toAppUser(ctx, r.rf, nil)
}

func (r *appUserRepository) findAppUserEntityByLoginID(ctx context.Context, organizationID *domain.OrganizationID, loginID string) (*appUserEntity, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindAppUserByLoginID")
	defer span.End()

	appUser := appUserEntity{}
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: organizationID}
	db := wrappedDB.WhereAppUser().Where(AppUserTableName+".login_id = ?", loginID).db
	if result := db.First(&appUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrAppUserNotFound
		}

		return nil, result.Error
	}

	return &appUser, nil
}

func (r *appUserRepository) FindOwnerByLoginID(ctx context.Context, operator service.SystemOwnerInterface, loginID string) (*service.Owner, error) {
	_, span := tracer.Start(ctx, "appUserRepository.FindOwnerByLoginID")
	defer span.End()

	appUser := appUserEntity{}
	wrappedDB := wrappedDB{dialect: r.dialect, db: r.db, organizationID: operator.OrganizationID()}
	db := wrappedDB.Table(AppUserTableName).Select(AppUserTableName+".*").
		WherePairOfUserAndGroup().
		WhereUserGroup().
		WhereAppUser().
		Where(AppUserTableName+".login_id = ?", loginID).
		Where(UserGroupTableName+".key_name = ? ", service.OwnerGroupKey).
		Joins("inner join " + PairOfUserAndGroupTableName + " on " + AppUserTableName + ".id = " + PairOfUserAndGroupTableName + ".app_user_id").
		Joins("inner join " + UserGroupTableName + " on " + PairOfUserAndGroupTableName + ".user_group_id = " + UserGroupTableName + ".id").
		db

	if result := db.First(&appUser); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrAppUserNotFound
		}

		return nil, result.Error
	}

	return appUser.toOwner(r.rf, nil)
}

func (r *appUserRepository) addAppUser(ctx context.Context, appUserEntity *appUserEntity) (*domain.AppUserID, error) {
	if result := r.db.Create(appUserEntity); result.Error != nil {
		return nil, liberrors.Errorf("db.Create. err: %w", libgateway.ConvertDuplicatedError(result.Error, service.ErrAppUserAlreadyExists))
	}

	appUserID, err := domain.NewAppUserID(appUserEntity.ID)
	if err != nil {
		return nil, err
	}

	return appUserID, nil
}

func (r *appUserRepository) AddAppUser(ctx context.Context, operator service.OwnerModelInterface, param service.AppUserAddParameterInterface) (*domain.AppUserID, error) {
	_, span := tracer.Start(ctx, "appUserRepository.AddAppUser")
	defer span.End()

	hashedPassword := ""
	if len(param.Password()) != 0 {
		hashedPasswordTmp, err := libgateway.HashPassword(param.Password())
		if err != nil {
			return nil, liberrors.Errorf("libgateway.HashPassword. err: %w", err)
		}

		hashedPassword = hashedPasswordTmp
	}

	appUserEntity := appUserEntity{
		BaseModelEntity: BaseModelEntity{
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: operator.OrganizationID().Int(),
		LoginID:        param.LoginID(),
		Username:       param.Username(),
		HashedPassword: hashedPassword,
	}

	appUserID, err := r.addAppUser(ctx, &appUserEntity)
	if err != nil {
		return nil, err
	}

	return appUserID, nil
}

func (r *appUserRepository) AddSystemOwner(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID) (*domain.AppUserID, error) {
	_, span := tracer.Start(ctx, "appUserRepository.AddSystemOwner")
	defer span.End()

	appUserEntity := appUserEntity{
		BaseModelEntity: BaseModelEntity{
			Version:   1,
			CreatedBy: operator.AppUserID().Int(),
			UpdatedBy: operator.AppUserID().Int(),
		},
		OrganizationID: organizationID.Int(),
		LoginID:        service.SystemOwnerLoginID,
		Username:       "SystemOwner",
	}

	appUserID, err := r.addAppUser(ctx, &appUserEntity)
	if err != nil {
		return nil, err
	}

	return appUserID, nil
}

func (r *appUserRepository) VerifyPassword(ctx context.Context, operator service.SystemAdminInterface, organizationID *domain.OrganizationID, loginID, password string) (bool, error) {
	appUserEntity, err := r.findAppUserEntityByLoginID(ctx, organizationID, loginID)
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
