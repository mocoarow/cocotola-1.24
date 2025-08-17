package controller_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
	"github.com/stretchr/testify/require"
)

// type Operator struct {
// 	organizationID *mbuserdomain.OrganizationID
// 	appUserID      *mbuserdomain.AppUserID
// }

// func newOperator(t *testing.T, organizationIDValue, appUserIDValue int) *Operator {
// 	t.Helper()
// 	return &Operator{
// 		organizationID: organizationID(t, organizationIDValue),
// 		appUserID:      appUserID(t, appUserIDValue),
// 	}
// }

// func (o *Operator) AppUserID() *mbuserdomain.AppUserID {
// 	return o.appUserID
// }
// func (o *Operator) OrganizationID() *mbuserdomain.OrganizationID {
// 	return o.organizationID
// }

// func organizationID(t *testing.T, organizationID int) *mbuserdomain.OrganizationID {
// 	t.Helper()
// 	id, err := mbuserdomain.NewOrganizationID(organizationID)
// 	require.NoError(t, err)
// 	return id
// }

// func appUserID(t *testing.T, appUserID int) *mbuserdomain.AppUserID {
// 	t.Helper()
// 	id, err := mbuserdomain.NewAppUserID(appUserID)
// 	require.NoError(t, err)
// 	return id
// }

func readBytes(t *testing.T, b *bytes.Buffer) []byte {
	t.Helper()
	respBytes, err := io.ReadAll(b)
	require.NoError(t, err)
	return respBytes
}

func parseJSON(t *testing.T, bytes []byte) interface{} {
	t.Helper()
	obj, err := oj.Parse(bytes)
	require.NoError(t, err)
	return obj
}

func parseExpr(t *testing.T, v string) jp.Expr {
	t.Helper()
	expr, err := jp.ParseString(v)
	require.NoError(t, err)
	return expr
}
