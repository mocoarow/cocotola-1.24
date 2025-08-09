package student

import (
	"context"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
)

type WorkbookQueryService interface {
	FindWorkbooks(ctx context.Context, operator service.OperatorInterface, param *libapi.WorkbookFindParameter) (*libapi.WorkbookFindResult, error)

	RetrieveWorkbookByID(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID) (*libapi.WorkbookRetrieveResult, error)
}

type WorkbookQueryUsecase struct {
	txManager            service.TransactionManager
	nonTxManager         service.TransactionManager
	workbookQuerySerivce WorkbookQueryService
}

func NewWorkbookQueryUsecase(txManager, nonTxManager service.TransactionManager, workbookQuerySerivce WorkbookQueryService) *WorkbookQueryUsecase {
	return &WorkbookQueryUsecase{
		txManager:            txManager,
		nonTxManager:         nonTxManager,
		workbookQuerySerivce: workbookQuerySerivce,
	}
}

func (u *WorkbookQueryUsecase) FindWorkbooks(ctx context.Context, operator service.OperatorInterface, param *libapi.WorkbookFindParameter) (*libapi.WorkbookFindResult, error) {
	workbooks, err := u.workbookQuerySerivce.FindWorkbooks(ctx, operator, param)
	if err != nil {
		return nil, err
	}

	return workbooks, nil

	// return &WorkbookFindResult{
	// 	TotalCount: 1,
	// 	Results: []*WorkbookFindWorkbookModel{
	// 		{
	// 			ID:   1,
	// 			Name: "test",
	// 		},
	// 	},
	// }, nil

	// var result domain.WorkbookSearchResult
	// fn := func(student service.Student) error {
	// 	condition, err := domain.NewWorkbookSearchCondition(DefaultPageNo, DefaultPageSize, []userD.SpaceID{})
	// 	if err != nil {
	// 		return rserrors.Errorf("service.NewWorkbookSearchCondition. err: %w", err)
	// 	}

	// 	tmpResult, err := student.FindWorkbooksFromPersonalSpace(ctx, condition)
	// 	if err != nil {
	// 		return rserrors.Errorf("student.FindWorkbooksFromPersonalSpace. err: %w", err)
	// 	}

	// 	result = tmpResult
	// 	return nil
	// }

	// if err := u.studentHandle(ctx, organizationID, operatorID, fn); err != nil {
	// 	return nil, err
	// }

	// return result, nil
}

func (u *WorkbookQueryUsecase) RetrieveWorkbookByID(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID) (*libapi.WorkbookRetrieveResult, error) {
	// TODO: check RBAC

	workbook, err := u.workbookQuerySerivce.RetrieveWorkbookByID(ctx, operator, workbookID)
	if err != nil {
		return nil, err
	}

	return workbook, nil
}
