package helpers

import (
	"mime/multipart"

	"github.com/golang-jwt/jwt/v5"
)

type JWTInterface interface {
	GenerateJWT(userID string, roleID string) map[string]any
	GenerateToken(userID string, roleID string) string
	ExtractToken(token *jwt.Token) any
	ValidateToken(token string, secret string) (*jwt.Token, error)
	RefereshJWT(refreshToken *jwt.Token) map[string]any
}

type HashInterface interface {
	HashPassword(password string) string
	CompareHash(password, hashed string) bool
}

type ValidationInterface interface {
	ValidateRequest(request any) []string
}

type OpenFileHeaderInterface interface {
	OpenFileHeader(fileHeader *multipart.FileHeader) multipart.File
}

type GeneratorInterface interface {
	GenerateRandomID() int
}

type OpenAIInterface interface {
	GetAppInformation(question string, qnaList map[string]string) (string, error)
	GetNewsContent(prompt string) (string, error)
}