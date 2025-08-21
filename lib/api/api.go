package api

type AppUserInfoResponse struct {
	AppUserID      int    `json:"appUserId"`
	OrganizationID int    `json:"organizationId"`
	LoginID        string `json:"loginId"`
	Username       string `json:"username"`
}

type PasswordAuthParameter struct {
	LoginID          string `json:"loginId" binding:"required"`
	Password         string `json:"password" binding:"required"`
	OrganizationName string `json:"organizationName" binding:"required"`
}

type AuthResponse struct {
	AccessToken  *string `json:"accessToken"`
	RefreshToken *string `json:"refreshToken"`
}

type RefreshTokenParameter struct {
	RefreshToken string `json:"refreshToken"`
}

type SynthesizeParameter struct {
	Lang5 string `json:"lang5" binding:"required,len=5"`
	Voice string `json:"voice"`
	Text  string `json:"text"`
}

type SynthesizeResponse struct {
	AudioContent           string `json:"audioContent"`
	AudioLengthMillisecond int    `json:"audioLengthMillisecond"`
}

// Find
type WorkbookFindParameter struct {
	PageNo   int
	PageSize int
}

type WorkbookFindWorkbookModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WorkbookFindResult struct {
	TotalCount int                          `json:"totalCount"`
	Results    []*WorkbookFindWorkbookModel `json:"results"`
}

// Retrieve
type WorkbookRetrieveResult struct {
	ID                  int                       `json:"id"`
	Version             int                       `json:"version"`
	Name                string                    `json:"name"`
	Lang2               string                    `json:"lang2" binding:"required"`
	Description         string                    `json:"description"`
	ProblemType         string                    `json:"problmeType"`
	EnglishSentences    *EnglishSentencesModel    `json:"englishSentences,omitempty"`
	EnglishConversation *EnglishConversationModel `json:"englishConversation,omitempty"`
}

type WorkbookAddParameter struct {
	Name        string `json:"name" binding:"required"`
	ProblemType string `json:"problemType" binding:"required"`
	Lang2       string `json:"lang2" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type WorkbookUpdateParameter struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
type EnglishConversationModel struct {
}

type EnglishSentenceModel struct {
	SrcLang2                  string `json:"srcLang2"`
	SrcAudioContent           string `json:"srcAudioContent"`
	SrcAudioLengthMillisecond int    `json:"srcAudioLengthMillisecond"`
	SrcText                   string `json:"srcText"`
	DstLang2                  string `json:"dstLang2"`
	DstAudioContent           string `json:"dstAudioContent"`
	DstAudioLengthMillisecond int    `json:"dstAudioLengthMillisecond"`
	DstText                   string `json:"dstText"`
}

type EnglishSentencesModel struct {
	Sentences []*EnglishSentenceModel `json:"sentences"`
}

type ActionObjectEffect struct {
	Action string `json:"action"`
	Object string `json:"object"`
	Effect string `json:"effect"`
}
type AddPolicyToUserParameter struct {
	OrganizationID           int                  `json:"organizationId"`
	AppUserID                int                  `json:"appUserId"`
	ListOfActionObjectEffect []ActionObjectEffect `json:"listOfActionObjectEffect"`
}

// Find
type DeckFindParameter struct {
	PageNo   int
	PageSize int
}

type DeckFindDeckModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DeckFindResult struct {
	TotalCount int                          `json:"totalCount"`
	Results    []*WorkbookFindWorkbookModel `json:"results"`
}

// Retrieve
type DeckRetrieveResult struct {
	ID          int    `json:"id"`
	Version     int    `json:"version"`
	Name        string `json:"name"`
	TemplateID  int    `json:"temlateId"`
	Lang2       string `json:"lang2" binding:"required"`
	Description string `json:"description"`
}

type DeckAddParameter struct {
	Name        string `json:"name" binding:"required"`
	TemplateID  int    `json:"temlateId"`
	Lang2       string `json:"lang2" binding:"required"`
	Description string `json:"description"`
}

type DeckUpdateParameter struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
