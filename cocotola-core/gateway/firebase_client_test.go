package gateway_test

import (
	"context"
	"testing"

	firebase "firebase.google.com/go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_VerifyFirebase(t *testing.T) {
	t.Skip()
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "cocotola-1-23-develop-24-11-02",
	})
	require.NoError(t, err)

	authClient, err := app.Auth(ctx)
	require.NoError(t, err)

	// token, err := authClient.VerifyIDToken(ctx, "")
	// require.NoError(t, err)
	// t.Logf("token: %v", token)
	// assert.Equal(t, "anonymous", token.Firebase.SignInProvider)

	userRecord, err := authClient.GetUserByEmail(ctx, "pecolynx@gmail.com")
	require.NoError(t, err)
	t.Logf("userRecord: %v", userRecord)
	assert.Fail(t, "not implemented yet")

	// token, err := authClient.VerifyIDToken(ctx, idToken)
	// require.NoError(t, err)
	// t.Logf("token: %v", token)
}
