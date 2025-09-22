package auth

import (
	libapimb "github.com/mocoarow/cocotola-1.24/lib/api/moonbeam"
)

type CallbackOnAddUserRequest struct {
	OrganizationID libapimb.OrganizationID `json:"organizationId" binding:"required"`
	UserID         libapimb.UserID         `json:"userId" binding:"required"`
}
