package space

type FindSpacesResponseSpace struct {
	ID   int    `json:"id" binding:"required"`
	Key  string `json:"key" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type FindSpacesResponse struct {
	Results []FindSpacesResponseSpace `json:"results"`
}
