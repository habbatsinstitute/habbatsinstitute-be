package main

import (
	"context"
	"fmt"
	"institute/config"
	"institute/features/auth"
	"institute/features/chatbot"
	"institute/features/course"
	"institute/features/ebook"
	"institute/features/gemini"
	"institute/features/item"
	"institute/features/news"
	"institute/features/product"
	realtimechat "institute/features/realtime_chat"
	"institute/features/user"
	"institute/helpers"
	"institute/middlewares"
	"institute/routes"
	"institute/utils"
	"institute/utils/websocket"
	"log"
	"net/http"

	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"

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

	cbh "institute/features/chatbot/handler"
	cbr "institute/features/chatbot/repository"
	cbu "institute/features/chatbot/usecase"

	ih "institute/features/item/handler"
	ir "institute/features/item/repository"
	iu "institute/features/item/usecase"

	rch "institute/features/realtime_chat/handler"
	rcr "institute/features/realtime_chat/repository"
	rcu "institute/features/realtime_chat/usecase"

	eh "institute/features/ebook/handler"
	er "institute/features/ebook/repository"
	eu "institute/features/ebook/usecase"

	ph "institute/features/product/handler"
	pr "institute/features/product/repository"
	pu "institute/features/product/usecase"

	gh "institute/features/gemini/handler"
	gr "institute/features/gemini/repository"
	gu "institute/features/gemini/usecase"
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
	routes.Chats(e, ChatHandler(cfg))
	routes.Ebooks(e, EbookHandler(), jwtService, *cfg)
	routes.Products(e, ProductHandler(cfg), jwtService, *cfg )
	routes.Geminis(e, GeminiHandler())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello anjay mabar!")
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

func EbookHandler() ebook.Handler{
	config := config.InitConfig()
	cdn := utils.CloudinaryInstance(*config)
	validator := helpers.NewValidationRequest()

	db := utils.InitDB()

	repo := er.New(db, cdn, config)
	nc := eu.New(repo, validator)
	return eh.New(nc)
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

func ItemHandler(cfg *config.ProgramConfig) item.Handler {
	db := utils.InitDB()


	repo := ir.New(db)
	uc := iu.New(repo)
	return ih.New(uc)
}

func ChatHandler(cfg *config.ProgramConfig) realtimechat.Handler {
	db := utils.InitDB()
	mongoDB := utils.ConnectMongo()
	collection := mongoDB.Collection("private_chat_histories")
	socket := websocket.NewServer()
	
	userRepo := rcr.New(db, collection)
	uc := rcu.New(socket, userRepo)
	return rch.New(uc)
}

func ProductHandler(cfg *config.ProgramConfig) product.Handler {
	db := utils.InitDB()
	cdn := utils.CloudinaryInstance(*cfg)
	validator := helpers.NewValidationRequest()

	repo := pr.New(db, cdn, cfg)
	uc := pu.New(repo, validator)
	return ph.New(uc)
}

func GeminiHandler() gemini.Handler {
	db := utils.InitDB()
	ctx := context.Background()
	// cfg := config.ProgramConfig{}.GEMINI_KEY

	genaiClient, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyASNEvvKq_TZSYI_Mz5N5ngPe7Lj3UWx18"))
    if err != nil {
        log.Fatalf("failed to initialize genai client: %v", err)
    }

	repo := gr.NewRepository(db)
	uc := gu.New(repo, genaiClient)
	return gh.New(uc)
}