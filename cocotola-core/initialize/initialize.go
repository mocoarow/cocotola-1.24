package initialize

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mblibgateway "github.com/mocoarow/cocotola-1.24/moonbeam/lib/gateway"

	libcontroller "github.com/mocoarow/cocotola-1.24/lib/controller/gin"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/config"
	controller "github.com/mocoarow/cocotola-1.24/cocotola-core/controller/gin"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

func Initialize(ctx context.Context, parent gin.IRouter, dialect mblibgateway.DialectRDBMS, driverName string, db *gorm.DB, logConfig *mblibconfig.LogConfig, coreConfig *config.CoreConfig) error {
	rff := func(ctx context.Context, db *gorm.DB) (service.RepositoryFactory, error) {
		return gateway.NewRepositoryFactory(ctx, dialect, driverName, db, time.UTC) // nolint:wrapcheck
	}
	rf, err := rff(ctx, db)
	if err != nil {
		return err
	}
	// init transaction manager
	txManager, err := mblibgateway.NewTransactionManagerT(db, rff)
	if err != nil {
		return err
	}
	// init non transaction manager
	nonTxManager, err := mblibgateway.NewNonTransactionManagerT(rf)
	if err != nil {
		return err
	}

	// init auth middleware
	bearerTokenAuthMiddleware, err := controller.InitBearerTokenAuthMiddleware(coreConfig.AuthAPIClient)
	if err != nil {
		return err
	}

	basicAuthMiddleware := gin.BasicAuth(gin.Accounts{
		coreConfig.CoreAPIServer.Username: coreConfig.CoreAPIServer.Password,
	})

	// init public and private router group functions
	publicRouterGroupFuncs := controller.GetPublicRouterGroupFuncs()

	bearerTokenPrivateRouterGroupFuncs, err := controller.GetBearerTokenPrivateRouterGroupFuncs(ctx, coreConfig, db, txManager, nonTxManager)
	if err != nil {
		return err
	}

	basicPrivateRouterGroupFuncs, err := controller.GetBasicPrivateRouterGroupFuncs(ctx, coreConfig, db, txManager, nonTxManager)
	if err != nil {
		return err
	}

	// api
	api := libcontroller.InitAPIRouterGroup(ctx, parent, domain.AppName, logConfig)

	// v1
	v1 := api.Group("v1")

	// public router
	libcontroller.InitPublicAPIRouterGroup(ctx, v1, publicRouterGroupFuncs)

	// private router
	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, bearerTokenAuthMiddleware, bearerTokenPrivateRouterGroupFuncs)

	libcontroller.InitPrivateAPIRouterGroup(ctx, v1, basicAuthMiddleware, basicPrivateRouterGroupFuncs)

	return nil
}

// const readHeaderTimeout = time.Duration(30) * time.Second

// type systemOwnerByOrganizationName struct {
// }

// func (s systemOwnerByOrganizationName) Get(ctx context.Context, rf service.RepositoryFactory, organizationName string) (*mbuserservice.SystemOwner, error) {
// 	mbrf, err := rf.NewmoonbeamRepositoryFactory(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, mbrf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return systemOwner, nil
// }

// func InitAppServer(ctx context.Context, rootRouterGroup gin.IRouter, corsConfig *mblibconfig.CORSConfig, debugConfig *libconfig.DebugConfig, appName string, authMiddleware gin.HandlerFunc, publicRouterGroupFuncs, privateRouterGroupFuncs []libcontroller.InitRouterGroupFunc) {
// 	// cors
// 	ginCorsConfig := mblibconfig.InitCORS(corsConfig)

// 	// root
// 	libcontroller.InitRootRouterGroup(ctx, rootRouterGroup, ginCorsConfig, debugConfig)

// 	InitApiServer(ctx, rootRouterGroup, appName, authMiddleware, publicRouterGroupFuncs, privateRouterGroupFuncs)
// }

// func InitApp1(ctx context.Context, txManager service.TransactionManager, workbookQueryService studentusecase.WorkbookQueryService) error {
// 	if err := txManager.Do(ctx, func(rf service.RepositoryFactory) error {
// 		// rf.NewWorkbookRepository(ctx)
// 		return nil
// 	}); err != nil {
// 		return err
// 	}

// 	workbookQueryService.RetrieveWorkbookByID(ctx)

// 	type Problem struct {
// 		Type       string            `json:"type"`
// 		Properties map[string]string `json:"properties"`
// 	}

// 	type Content struct {
// 		Problems []*Problem `json:"problems"`
// 	}

// 	x := Content{
// 		Problems: []*Problem{
// 			{
// 				Type: "text",
// 				Properties: map[string]string{
// 					"srcLang":         "ja",
// 					"srcAudioContent": audioContentJa1,
// 					"srcAudioLength":  strconv.Itoa(audioLengthJa1),
// 					"srcText":         "こんにちは",
// 					"dstLang":         "en",
// 					"dstAudioContent": audioContentEn1,
// 					"dstAudioLength":  strconv.Itoa(audioLengthEn1),
// 					"dstText":         "Hello",
// 				},
// 			},
// 			{
// 				Type: "text",
// 				Properties: map[string]string{
// 					"srcLang":         "ja",
// 					"srcAudioContent": audioContentJa2,
// 					"srcAudioLength":  strconv.Itoa(audioLengthJa2),
// 					"srcText":         "さようなら",
// 					"dstLang":         "en",
// 					"dstAudioContent": audioContentEn2,
// 					"dstAudioLength":  strconv.Itoa(audioLengthEn2),
// 					"dstText":         "Goodbye",
// 				},
// 			},
// 		},
// 	}

// 	_, err := json.Marshal(x)
// 	if err != nil {
// 		return err
// 	}

// 	// fmt.Println(jsonBytes)

// 	return nil
// }

// func systemOwnerAction(ctx context.Context, organizationName string, txManager service.TransactionManager, fn func(context.Context, *mbuserservice.SystemOwner) error) error {
// 	return txManager.Do(ctx, func(rf service.RepositoryFactory) error {
// 		mbrf, err := rf.NewmoonbeamRepositoryFactory(ctx)
// 		if err != nil {
// 			return mbliberrors.Errorf(". err: %w", err)
// 		}

// 		systemAdmin, err := mbuserservice.NewSystemAdmin(ctx, mbrf)
// 		if err != nil {
// 			return mbliberrors.Errorf(". err: %w", err)
// 		}
// 		systemOwner, err := systemAdmin.FindSystemOwnerByOrganizationName(ctx, organizationName)
// 		if err != nil {
// 			return mbliberrors.Errorf(". err: %w", err)
// 		}

// 		return fn(ctx, systemOwner)
// 	})
// }
