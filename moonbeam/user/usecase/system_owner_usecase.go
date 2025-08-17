package usecase

type SystemOwnerUsecase interface {
	Init() error
}

type systemOwnerUsecase struct {
}

func NewSystemOwnerUsecase() SystemOwnerUsecase {
	return &systemOwnerUsecase{}
}

func (u *systemOwnerUsecase) Init() error {
	return nil
}
