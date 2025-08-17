package service

import (
	"context"
	"errors"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
)

var ErrWorkbookAlreadyExists = errors.New("workbook already exists")

var ErrWorkbookNotFound = errors.New("workbook not found")

type WorkbookAddParameter struct {
	Name        string
	ProblemType string
	Lang2       string
	Description string
	Content     string
}

type WorkbookUpdateParameter struct {
	Name        string
	Description string
	Content     string
}

type OperatorInterface interface {
	AppUserID() *mbuserdomain.AppUserID
	OrganizationID() *mbuserdomain.OrganizationID
	// LoginID() string
	// Username() string
}

type WorkbookRepository interface {
	AddWorkbook(ctx context.Context, operator OperatorInterface, param *WorkbookAddParameter) (*domain.WorkbookID, error)

	UpdateWorkbook(ctx context.Context, operator OperatorInterface, workbookID *domain.WorkbookID, version int, param *WorkbookUpdateParameter) error
}
