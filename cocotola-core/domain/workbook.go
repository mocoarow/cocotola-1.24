package domain

type WorkbookID struct {
	Value int `validate:"required,gte=1"`
}

func NewWorkbookID(value int) (*WorkbookID, error) {
	return &WorkbookID{
		Value: value,
	}, nil
}

func (v *WorkbookID) Int() int {
	return v.Value
}
func (v *WorkbookID) IsWorkbookID() bool {
	return true
}
