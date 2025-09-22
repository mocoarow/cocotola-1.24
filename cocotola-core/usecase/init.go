package usecase

import (
	"go.opentelemetry.io/otel"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type Operator struct {
	userID         *mbuserdomain.UserID
	organizationID *mbuserdomain.OrganizationID
}

func (o *Operator) GetUserID() *mbuserdomain.UserID {
	return o.userID
}
func (o *Operator) GetOrganizationID() *mbuserdomain.OrganizationID {
	return o.organizationID
}

var (
	tracer = otel.Tracer("github.com/mocoarow/cocotola-1.24/cocotola-core/usecase")
)
