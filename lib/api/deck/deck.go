package card

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
