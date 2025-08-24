//go:build scenario

package scenario_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/k1LoW/runn"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	ctx := context.Background()

	newUserLoginID, err := uuid.NewRandom()
	require.NoError(t, err)
	newUserUsername, err := uuid.NewRandom()
	require.NoError(t, err)
	newUserPassword, err := uuid.NewRandom()
	require.NoError(t, err)

	opts := []runn.Option{
		runn.Scopes("run:exec"),
		runn.T(t),
		runn.Runner("req", "http://localhost:8000"),
		runn.Var("NEW_USER_LOGIN_ID", newUserLoginID.String()),
		runn.Var("NEW_USER_USERNAME", newUserUsername.String()),
		runn.Var("NEW_USER_PASSWORD", newUserPassword.String()),
	}
	o, err := runn.Load("runntest.yml", opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(ctx); err != nil {
		t.Fatal(err)
	}
}
