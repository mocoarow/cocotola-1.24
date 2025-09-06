package card

// FindCardsResponse
type FindCardsResponse struct {
	TotalCount int                     `json:"totalCount"`
	Results    []FindCardsResponseCard `json:"results"`
}

type FindCardsResponseCard struct {
	ID         int    `json:"id" binding:"required"`
	Version    int    `json:"version" binding:"required"`
	TemplateID int    `json:"templateId" binding:"required"`
	Content    string `json:"name" binding:"required"`
}
