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

func registerUser(ctx context.Context, systemToken libdomain.SystemToken, rf service.RepositoryFactory, organizationID *mbuserdomain.OrganizationID, loginID string, createUserParameterFunc func() (*mbuserservice.AddUserParameter, error)) (*mbuserdomain.OrganizationModel, *mbuserdomain.UserModel, error) {
	action, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
		service.WithOrganizationByID(organizationID),
	)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("NewSystemOwnerAction: %w", err)
	}

	if _, err = action.SystemOwner.FindUserByLoginID(ctx, loginID); err == nil {
		return nil, nil, mbuserservice.ErrUserAlreadyExists
	} else if !errors.Is(err, mbuserservice.ErrUserNotFound) {
		return nil, nil, mbliberrors.Errorf("systemOwner.FindUserByLoginID. err: %w", err)
	}

	appUser, err := registerUserWithSystemOwnerAction(ctx, action, createUserParameterFunc)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("find or register app user: %w", err)
	}

	return action.Organization.OrganizationModel, appUser, nil
}

func findOrRegisterUser(ctx context.Context, systemToken libdomain.SystemToken, rf service.RepositoryFactory, organizationID *mbuserdomain.OrganizationID, loginID string, createUserParameterFunc func() (*mbuserservice.AddUserParameter, error)) (*mbuserdomain.OrganizationModel, *mbuserdomain.UserModel, error) {
	action, err := service.NewSystemOwnerAction(ctx, systemToken, rf,
		service.WithOrganizationByID(organizationID),
	)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("NewSystemOwnerAction: %w", err)
	}

	appUser1, err := action.SystemOwner.FindUserByLoginID(ctx, loginID)
	if err == nil {
		return action.Organization.OrganizationModel, appUser1.UserModel, nil
	} else if !errors.Is(err, mbuserservice.ErrUserNotFound) {
		return nil, nil, mbliberrors.Errorf("systemOwner.FindUserByLoginID. err: %w", err)
	}

	appUser, err := registerUserWithSystemOwnerAction(ctx, action, createUserParameterFunc)
	if err != nil {
		return nil, nil, mbliberrors.Errorf("find or register app user: %w", err)
	}

	return action.Organization.OrganizationModel, appUser, nil
}

func registerUserWithSystemOwnerAction(ctx context.Context, systemOwnerAction *service.SystemOwnerAction, createUserParameterFunc func() (*mbuserservice.AddUserParameter, error)) (*mbuserdomain.UserModel, error) {
	parameter, err := createUserParameterFunc()
	if err != nil {
		return nil, mbliberrors.Errorf("invalid UserAddParameter. err: %w", err)
	}

	studentID, err := systemOwnerAction.SystemOwner.AddUser(ctx, parameter)
	if err != nil {
		return nil, mbliberrors.Errorf("failed to AddStudent. err: %w", err)
	}

	appUser2, err := systemOwnerAction.SystemOwner.FindUserByID(ctx, studentID)
	if err != nil {
		return nil, mbliberrors.Errorf("failed to FindStudentByID. err: %w", err)
	}

	return appUser2.UserModel, nil
}
