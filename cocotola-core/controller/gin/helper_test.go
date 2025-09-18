package controller_test

// type Operator struct {
// 	organizationID *mbuserdomain.OrganizationID
// 	userID      *mbuserdomain.UserID
// }

// func newOperator(t *testing.T, organizationIDValue, userIDValue int) *Operator {
// 	t.Helper()
// 	return &Operator{
// 		organizationID: organizationID(t, organizationIDValue),
// 		userID:      userID(t, userIDValue),
// 	}
// }

// func (o *Operator) UserID() *mbuserdomain.UserID {
// 	return o.userID
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

// func userID(t *testing.T, userID int) *mbuserdomain.UserID {
// 	t.Helper()
// 	id, err := mbuserdomain.NewUserID(userID)
// 	require.NoError(t, err)
// 	return id
// }

/*
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
*/
