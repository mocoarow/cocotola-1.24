package auth

type ActionObjectEffect struct {
	Action string `json:"action"`
	Object string `json:"object"`
	Effect string `json:"effect"`
}
type ActionObject struct {
	Action string `json:"action"`
	Object string `json:"object"`
}
type AddPolicyToUserParameter struct {
	OrganizationID           int                  `json:"organizationId"`
	AppUserID                int                  `json:"appUserId"`
	ListOfActionObjectEffect []ActionObjectEffect `json:"listOfActionObjectEffect"`
}

type AuthorizeRequest struct {
	OrganizationID int    `json:"organizationId"`
	AppUserID      int    `json:"appUserId"`
	Action         string `json:"action"`
	Object         string `json:"object"`
}

type AuthorizeResponse struct {
	Authorized bool `json:"authorized"`
}
