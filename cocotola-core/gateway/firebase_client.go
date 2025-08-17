package gateway

// func VerifyFirebase(ctx context.Context, clientOption option.ClientOption, idToken string) error {
// 	app, err := firebase.NewApp(ctx, nil, clientOption)
// 	if err != nil {
// 		return err
// 	}
// 	authClient, err := app.Auth(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	userRecord, err:=authClient.GetUserByEmail(ctx, "pecolynx@gmail.com")

// 	token, err := authClient.VerifyIDToken(ctx, idToken)
// 	if err != nil {
// 		// エラー処理
// 	}
// 	return nil
// }
