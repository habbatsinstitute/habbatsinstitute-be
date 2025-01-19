package dtos

type InformationBodyRequest struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type InformationResponse struct {
	ID       int  `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type ChatRequestPayload struct {
	Question string `json:"question" binding:"required"`
}

type ChatResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}