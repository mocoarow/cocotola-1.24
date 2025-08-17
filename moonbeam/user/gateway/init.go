package gateway

import (
	"go.opentelemetry.io/otel"
)

var (
	tracer = otel.Tracer("github.com/mocoarow/cocotola-1.24/moonbeam/user/gateway")

	OrganizationTableName       = "mb_organization"
	AppUserTableName            = "mb_app_user"
	PairOfUserAndGroupTableName = "mb_user_n_group"
	UserGroupTableName          = "mb_user_group"

	// SystemStudentLoginID = "system-student"
	// GuestLoginID         = "guest"

	// AdministratorRole = "Administrator"
	// ManagerRole       = "Manager"
	// UserRole          = "User"
	// GuestRole         = "Guest"
	// UnknownRole       = "Unknown"
)
