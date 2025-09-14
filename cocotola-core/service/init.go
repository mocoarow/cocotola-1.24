package service

import "github.com/mocoarow/cocotola-1.24/cocotola-core/domain"

var (
	TEMPLATE_ID_ENGLISH_BLANK *domain.TemplateID
)

func init() {
	templateIDEnglishBlank, err := domain.NewTemplateID(1)
	if err != nil {
		panic(err)
	}
	TEMPLATE_ID_ENGLISH_BLANK = templateIDEnglishBlank
}
