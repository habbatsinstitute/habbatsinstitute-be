package main

import (
	"fmt"
	"institute/config"
	"institute/features/auth"
	"institute/features/course"
	"institute/features/news"
	"institute/features/user"
	"institute/helpers"
	"institute/middlewares"
	"institute/routes"
	"institute/utils"
	"net/http"

	"github.com/labstack/echo/v4"

	ah "institute/features/auth/handler"
	ar "institute/features/auth/repository"
	au "institute/features/auth/usecase"

	ch "institute/features/course/handler"
	cr "institute/features/course/repository"
	cu "institute/features/course/usecase"

	uh "institute/features/user/handler"
	ur "institute/features/user/repository"
	uu "institute/features/user/usecase"

	nh "institute/features/news/handler"
	nr "institute/features/news/repository"
	nu "institute/features/news/usecase"

	"institute/features/chatbot"
	cbh "institute/features/chatbot/handler"
	cbr "institute/features/chatbot/repository"
	cbu "institute/features/chatbot/usecase"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	jwtService := helpers.NewJWT(*cfg)
	
	middlewares.LogMiddlewares(e)
	routes.Auth(e, AuthHandler(), jwtService, *cfg)
	routes.Courses(e, CourseHandler(), jwtService, *cfg)
	routes.Users(e, UserHandler(), jwtService, *cfg)
	routes.Newss(e, NewsHandler(), jwtService, *cfg)
	routes.Chatbots(e, ChatbotHandler(cfg), jwtService, *cfg)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!")
	})
	e.Start(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func AuthHandler() auth.Handler{
	config := config.InitConfig()

	db := utils.InitDB()
	jwt := helpers.NewJWT(*config)
	hash := helpers.NewHash()
	validation := helpers.NewValidationRequest()

	repo := ar.New(db)
	ac := au.New(repo, jwt, hash, validation)
	return ah.New(ac)
}

func CourseHandler() course.Handler {
	config := config.InitConfig()
	cdn := utils.CloudinaryInstance(*config)
	jwt := helpers.NewJWT(*config)
	validator := helpers.NewValidationRequest()

	db := utils.InitDB()

	repo := cr.New(db, cdn, config)
	cc :=	cu.New(repo, jwt, validator)
	return ch.New(cc)
}

func UserHandler() user.Handler {
	config := config.InitConfig()
	jwt := helpers.NewJWT(*config)
	hash := helpers.NewHash()

	db := utils.InitDB()

	repo := ur.New(db)
	uc := uu.New(repo, jwt, hash)
	return uh.New(uc)
}

func NewsHandler() news.Handler {
	config := config.InitConfig()
	cdn := utils.CloudinaryInstance(*config)
	validator := helpers.NewValidationRequest()

	db := utils.InitDB()

	repo := nr.New(db, cdn, config)
	nc := nu.New(repo, validator)
	return nh.New(nc)

}

func ChatbotHandler(cfg *config.ProgramConfig) chatbot.Handler {
	db := utils.InitDB()
	mongoDB := utils.ConnectMongo()
	collection := mongoDB.Collection("chatbot_histories")

	validation := helpers.NewValidationRequest()
	openAI := helpers.NewOpenAI(cfg.OPENAI_KEY)

	repo := cbr.New(db, collection)
	uc := cbu.New(repo, validation, openAI)
	return cbh.New(uc)
}