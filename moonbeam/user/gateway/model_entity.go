package gateway

import (
	"time"

	libdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	liberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"
)

type BaseModelEntity struct {
	Version   int
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy int
	UpdatedBy int
}

func (e *BaseModelEntity) toBaseModel() (*libdomain.BaseModel, error) {
	model, err := libdomain.NewBaseModel(e.Version, e.CreatedAt, e.UpdatedAt, e.CreatedBy, e.UpdatedBy)
	if err != nil {
		return nil, liberrors.Errorf("libdomain.NewBaseModel. err: %w", err)
	}

	return model, nil
}

type JunctionModelEntity struct {
	CreatedAt time.Time
	CreatedBy int
}

// type JunctionModelEntity struct {
// 	CreatedAt time.Time
// 	CreatedBy uint
// }

// // func (e *junctionModelEntity) toModel() (domain.Model, error) {
// // 	return domain.NewModel(e.ID, e.Version, e.CreatedAt, e.UpdatedAt, e.CreatedBy, e.UpdatedBy)
// // }
