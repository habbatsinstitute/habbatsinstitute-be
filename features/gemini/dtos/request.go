package dtos

type InformationBodyRequest struct {
	Question string `json:"message" form:"message" binding:"required"`
	Answer   string `json:"reply" binding:"required"`
}

type InformationResponse struct {
	ID       int  `json:"id"`
	Question string `json:"message" form:"message"`
	Answer   string `json:"reply"`
}

type ChatRequestPayload struct {
	Question string `json:"message" form:"message" binding:"required"`
}

type ChatResponse struct {
	Question string `json:"message" form:"message"`
	Answer   string `json:"reply"`
}