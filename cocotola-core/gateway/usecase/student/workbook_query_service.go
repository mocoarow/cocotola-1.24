package student

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/gateway"
	"github.com/mocoarow/cocotola-1.24/cocotola-core/service"
	studentusecase "github.com/mocoarow/cocotola-1.24/cocotola-core/usecase/student"

	libapi "github.com/mocoarow/cocotola-1.24/lib/api"
)

type workbookQueryService struct {
	db *gorm.DB
}

func NewWorkbookQueryService(db *gorm.DB) studentusecase.WorkbookQueryService {
	return &workbookQueryService{
		db: db,
	}
}

func (s *workbookQueryService) FindWorkbooks(ctx context.Context, operator service.OperatorInterface, param *libapi.WorkbookFindParameter) (*libapi.WorkbookFindResult, error) {
	return nil, nil
}

func (s *workbookQueryService) RetrieveWorkbookByID(ctx context.Context, operator service.OperatorInterface, workbookID *domain.WorkbookID) (*libapi.WorkbookRetrieveResult, error) {
	workbookE := gateway.WorkbookEntity{}
	if result := s.db.Where("workbook.id = ?", workbookID.Int()).First(&workbookE); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, service.ErrWorkbookNotFound
		}
		return nil, result.Error
	}

	return workbookE.ToModel()
}
