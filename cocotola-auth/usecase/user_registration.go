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
	action, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
		service.WithOrganizationByID(organizationID),
	)
	if err != nil {
		return nil, nil, err
	}

	if _, err = action.SystemOwner.FindAppUserByLoginID(ctx, loginID); err == nil {
		return nil, nil, mbuserservice.ErrAppUserAlreadyExists
	} else if !errors.Is(err, mbuserservice.ErrAppUserNotFound) {
		return nil, nil, mbliberrors.Errorf("systemOwner.FindAppUserByLoginID. err: %w", err)
	}

	appUser, err := registerAppUserWithSystemOwnerAction(ctx, action, createAppUserParameterFunc)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("find or register app user: %w", err)
	}

	return action.Organization.OrganizationModel, appUser, nil
}

func findOrRegisterAppUser(ctx context.Context, systemToken libdomain.SystemToken, rf service.RepositoryFactory, organizationID *mbuserdomain.OrganizationID, loginID string, createAppUserParameterFunc func() (*mbuserservice.AppUserAddParameter, error)) (*mbuserdomain.OrganizationModel, *mbuserdomain.AppUserModel, error) {
	action, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
		service.WithOrganizationByID(organizationID),
	)
	if err != nil {
		return nil, nil, err
	}

	appUser1, err := action.SystemOwner.FindAppUserByLoginID(ctx, loginID)
	if err == nil {
		return action.Organization.OrganizationModel, appUser1.AppUserModel, nil
	} else if !errors.Is(err, mbuserservice.ErrAppUserNotFound) {
		return nil, nil, mbliberrors.Errorf("systemOwner.FindAppUserByLoginID. err: %w", err)
	}

	appUser, err := registerAppUserWithSystemOwnerAction(ctx, action, createAppUserParameterFunc)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("find or register app user: %w", err)
	}

	return action.Organization.OrganizationModel, appUser, nil
}

func registerAppUserWithSystemOwnerAction(ctx context.Context, systemOwnerAction *service.SystemOwnerAction, createAppUserParameterFunc func() (*mbuserservice.AppUserAddParameter, error)) (*mbuserdomain.AppUserModel, error) {
	parameter, err := createAppUserParameterFunc()
	if err != nil {
		return nil, mbliberrors.Errorf("invalid AppUserAddParameter. err: %w", err)
	}

	studentID, err := systemOwnerAction.SystemOwner.AddAppUser(ctx, parameter)
	if err != nil {
		return nil, mbliberrors.Errorf("failed to AddStudent. err: %w", err)
	}

	appUser2, err := systemOwnerAction.SystemOwner.FindAppUserByID(ctx, studentID)
	if err != nil {
		return nil, mbliberrors.Errorf("failed to FindStudentByID. err: %w", err)
	}

	return appUser2.AppUserModel, nil
}
