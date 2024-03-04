package dtos

type Request struct {
	Text        string `json:"text"`
	Blob        string `json:"blob"`
	RecipientID int    `json:"recipient_id"`
}