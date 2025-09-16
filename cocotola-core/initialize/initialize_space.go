package initialize

// func addPublicSpaceIfNotExists(ctx context.Context, rf service.RepositoryFactory, operator mbuserservice.OperatorInterface) (*domain.SpaceID, error) {
// 	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitPublicSpace"))

// 	spaceRepo, err := rf.NewSpaceRepository(ctx)
// 	if err != nil {
// 		return nil, mbliberrors.Errorf("NewSpaceRepository: %w", err)
// 	}

// 	// check default-public space
// 	space, err := spaceRepo.FindPublicSpaceByKey(ctx, "default-public")
// 	if err == nil {
// 		logger.Info("default-public space already exists", slog.Int("spaceID", space.SpaceID.Int()))
// 		return space.SpaceID, nil
// 	}

// 	if errors.Is(err, service.ErrSpaceNotFound) {
// 		spaceID, err := spaceRepo.AddSpace(ctx, operator, &service.SpaceAddParameter{
// 			Name:     "Default Public Space",
// 			Key:      "default-public",
// 			IsPublic: true,
// 		})
// 		if err != nil {
// 			return nil, mbliberrors.Errorf("AddSpace: %w", err)
// 		}
// 		logger.Info("default-public space created", slog.Int("spaceID", spaceID.Int()))

// 		return spaceID, nil
// 	}

// 	return nil, mbliberrors.Errorf("FindPublicSpaceByKey: %w", err)
// }

// func initApp1(ctx context.Context, txManager service.TransactionManager, authInitParam *AuthInitParameter) (*domain.SpaceID, error) {
// 	logger := slog.Default().With(slog.String(mbliblog.LoggerNameKey, domain.AppName+"InitApp1"))

// 	operator := &operator{
// 		organizationID: authInitParam.OrganizationID,
// 		userID:      mbuserservice.SystemAdminID,
// 	}
// 	fn := func(rf service.RepositoryFactory) (*domain.SpaceID, error) {
// 		spaceRepo, err := rf.NewSpaceRepository(ctx)
// 		if err != nil {
// 			return nil, mbliberrors.Errorf("NewSpaceRepository: %w", err)
// 		}

// 		// check default-public space
// 		spaceID, err := addPublicSpaceIfNotExists(ctx, rf, operator)
// 		if err != nil {
// 			return nil, mbliberrors.Errorf("addPublicSpaceIfNotExists: %w", err)
// 		}

// 		// guest user can access to the space

// 	}

// 	spaceID, err := mblibservice.Do1(ctx, txManager, fn)
// 	if err != nil {
// 		return nil, err //nolint:wrapcheck
// 	}

// 	return spaceID, nil
// }
