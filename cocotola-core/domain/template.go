package domain

import (
	"fmt"

	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

type TemplateID struct {
	Value int `validate:"required,gte=1"`
}

func NewTemplateID(value int) (*TemplateID, error) {
	return &TemplateID{
		Value: value,
	}, nil
}

func (v *TemplateID) Int() int {
	return v.Value
}
func (v *TemplateID) IsTemplateID() bool {
	return true
}
func (v *TemplateID) GetRBACObject() mbuserdomain.RBACObject {
	return mbuserdomain.NewRBACObject("template:" + fmt.Sprint(v.Value))
}
