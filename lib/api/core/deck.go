package core

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

// // Find
// type DeckFindParameter struct {
// 	PageNo   int
// 	PageSize int
// }

// type DeckFindDeckModel struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

type UpdateDeckRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// Retrieve
type DeckRetrieveResult struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	Name        string `json:"name" binding:"required"`
	TemplateID  int    `json:"templateId" binding:"required"`
	Lang2       string `json:"lang2" binding:"required"`
	Description string `json:"description"`
}

// AddDeckRequest
type AddDeckRequest struct {
	SpaceID     int    `json:"spaceId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	TemplateID  int    `json:"templateId" binding:"required"`
	Lang2       string `json:"lang2" binding:"required"`
	Description string `json:"description"`
}
