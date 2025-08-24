package usecase

import (
	"context"
	"log/slog"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbliblog "github.com/mocoarow/cocotola-1.24/moonbeam/lib/log"
	mblibservice "github.com/mocoarow/cocotola-1.24/moonbeam/lib/service"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type Callback struct {
	txManager    service.TransactionManager
	nonTxManager service.TransactionManager
	rbacClient   service.CocotolaRBACClient
	logger       *slog.Logger
}

func NewCallback(txManager, nonTxManager service.TransactionManager, rbacClient service.CocotolaRBACClient) *Callback {
	return &Callback{
		txManager:    txManager,
		nonTxManager: nonTxManager,
		rbacClient:   rbacClient,
		logger:       slog.Default().With(slog.String(mbliblog.LoggerNameKey, "CallbackUsecase"))}
}

func (u *Callback) OnAddAppUser(ctx context.Context, organizationID *mbuserdomain.OrganizationID, appUserID *mbuserdomain.AppUserID) error {
	u.logger.InfoContext(ctx, "OnAddAppUser", slog.Int("app_user_id", appUserID.Int()))
	if err := mblibservice.Do0(ctx, u.nonTxManager, func(rf service.RepositoryFactory) error {
		spaceRepo, err := rf.NewSpaceRepository(ctx)
		if err != nil {
			return err
		}
		pairOfUserAndSpaceRep, err := rf.NewPairOfUserAndSpaceRepository(ctx)
		if err != nil {
			return err
		}

		operator := Operator{
			organizationID: organizationID,
			appUserID:      appUserID,
		}

		param := service.SpaceAddParameter{
			Key:  "private",
			Name: "Private",
		}

		spaceID, err := spaceRepo.AddSpace(ctx, &operator, &param)
		if err != nil {
			return err
		}

		object := spaceID.GetRBACObject()

		if err := u.rbacClient.AddPolicyToUser(ctx, &libapi.AddPolicyToUserParameter{
			OrganizationID: operator.OrganizationID().Int(),
			AppUserID:      operator.AppUserID().Int(),
			ListOfActionObjectEffect: []libapi.ActionObjectEffect{
				{
					Action: mbuserdomain.NewRBACAction("CreateDeck").Action(),
					Object: object.Object(),
					Effect: mbuserservice.RBACAllowEffect.Effect(),
				},
			},
		}); err != nil {
			return mbliberrors.Errorf("add policy to user. space(%d): %w", spaceID.Int(), err)
		}

		if err := pairOfUserAndSpaceRep.AddPairOfUserAndSpace(ctx, &operator, appUserID, spaceID); err != nil {
			return err
		}

		u.logger.InfoContext(ctx, "OnAddAppUser: AddSpace", slog.Int("space_id", spaceID.Int()))

		return nil
	}); err != nil {
		return err
	}
	return nil
}
