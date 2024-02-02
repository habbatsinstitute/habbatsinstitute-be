package chatbot

import (
	"institute/features/chatbot/dtos"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	SaveChat(questionNReply QuestionAndReply, userID int) error 
	SelectUserByID(userID int) (*User, error)
	SelectByUserID(userID int) (*ChatHistory, error)
	DeleteByUserID(userID int)error
	ReadQuestionNPrompts() (map[string]string, error)
	GetTimeNow() time.Time
}

type Usecase interface {
	FindAllChat(userID int) []dtos.ResChatReply
	SetContentForNews(input dtos.InputMessage) (*dtos.ResNewsContent, []string, error)
	SetReplyMessage(requestMessage dtos.InputMessage, userID int) (*dtos.ResChatReply, []string, error)
	ClearHistory(userID int) error
}

type Handler interface {
	GetChatHistory() echo.HandlerFunc
	GetNewsContentGeneration() echo.HandlerFunc
	SendQuestion() echo.HandlerFunc
	DeleteChatHistory() echo.HandlerFunc
}