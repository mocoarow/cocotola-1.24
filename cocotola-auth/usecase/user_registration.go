package usecase

import (
	"context"
	"errors"

	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
	mbuserservice "github.com/mocoarow/cocotola-1.24/moonbeam/user/service"

	libdomain "github.com/mocoarow/cocotola-1.24/lib/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-auth/service"
)

func registerAppUser(ctx context.Context, systemToken libdomain.SystemToken, rf service.RepositoryFactory, organizationID *mbuserdomain.OrganizationID, loginID string, createAppUserParameterFunc func() (*mbuserservice.AppUserAddParameter, error)) (*mbuserdomain.OrganizationModel, *mbuserdomain.AppUserModel, error) {
	action, err := service.NewOrganizationAction(ctx, systemToken, rf,
		service.WithOrganizationRepository(),
		service.WithOrganizationByID(organizationID),
		service.WithAppUserRepository(),
	)
	if err != nil {
		return nil, nil, err
	}

	findOrRegisterAppUser := func() (*mbuserdomain.AppUserModel, error) {
		appUser1, err := action.SystemOwner.FindAppUserByLoginID(ctx, loginID)
		if err == nil {
			return appUser1.AppUserModel, nil
		}

		if !errors.Is(err, mbuserservice.ErrAppUserNotFound) {
			return nil, mbliberrors.Errorf("systemOwner.FindAppUserByLoginID. err: %w", err)
		}

		parameter, err := createAppUserParameterFunc()
		if err != nil {
			return nil, mbliberrors.Errorf("invalid AppUserAddParameter. err: %w", err)
		}

		studentID, err := action.SystemOwner.AddAppUser(ctx, parameter)
		if err != nil {
			return nil, mbliberrors.Errorf("failed to AddStudent. err: %w", err)
		}

		appUser2, err := action.SystemOwner.FindAppUserByID(ctx, studentID)
		if err != nil {
			return nil, mbliberrors.Errorf("failed to FindStudentByID. err: %w", err)
		}

		return appUser2.AppUserModel, nil
	}

	appUser, err := findOrRegisterAppUser()
	if err != nil {
		return nil, nil, mbliberrors.Errorf("find or register app user: %w", err)
	}

	return action.Organization.OrganizationModel, appUser, nil
}
