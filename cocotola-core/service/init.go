package service

import "github.com/mocoarow/cocotola-1.24/cocotola-core/domain"

var (
	TemplateIDEnglishBlank *domain.TemplateID
)

func init() {
	templateIDEnglishBlank, err := domain.NewTemplateID(1)
	if err != nil {
		panic(err)
	}
	TemplateIDEnglishBlank = templateIDEnglishBlank
}
