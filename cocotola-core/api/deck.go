package api

import (
	"encoding/json"

	"github.com/mocoarow/cocotola-1.24/cocotola-core/domain"
	libapimb "github.com/mocoarow/cocotola-1.24/lib/api/moonbeam"
)

type FindDecksRequest struct {
	SpaceID []int `form:"spaceId"`
}

// FindDecksResponse
type FindDecksResponse struct {
	TotalCount int                     `json:"totalCount"`
	Results    []FindDecksResponseDeck `json:"results"`
}

type FindDecksResponseDeck struct {
	ID          int    `json:"id" binding:"required"`
	Version     int    `json:"version" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Lang2       string `json:"lang2" binding:"required"`
	TemplateID  int    `json:"templateId" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type AddDeckRequest struct {
	SpaceID     libapimb.SpaceID `json:"spaceId" binding:"required"`
	Name        string           `json:"name" binding:"required"`
	TemplateID  TemplateID       `json:"templateId" binding:"required"`
	Lang2       libapimb.Lang2   `json:"lang2" binding:"required"`
	Description string           `json:"description"`
}

type TemplateID struct {
	Value *domain.TemplateID `validate:"required,gte=1"`
}

func (m *TemplateID) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	value, err := domain.NewTemplateID(v)
	if err != nil {
		return err
	}
	m.Value = value
	return nil
}

func (m TemplateID) MarshalJSON() ([]byte, error) {
	if m.Value == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(m.Value.Int())
}
