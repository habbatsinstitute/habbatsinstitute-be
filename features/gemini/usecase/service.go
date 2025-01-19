package usecase

import (
	"institute/features/gemini"
	"institute/features/gemini/dtos"
	"institute/helpers"

	"github.com/google/generative-ai-go/genai"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type service struct {
	model gemini.Repository
	genai *genai.Client
}

func New(model gemini.Repository, genai *genai.Client) gemini.Usecase {
	return &service {
		model: model,
		genai: genai,
	}
}

func (uc *service) InsertInformation(ctx echo.Context, payload dtos.InformationBodyRequest) (information dtos.InformationResponse, err error) {

	result, err := uc.model.InsertInformation(ctx, payload)
	if err != nil {
		return
	}
	information = dtos.InformationResponse{
		ID:       result,
		Answer:   payload.Answer,
		Question: payload.Question,
	}
	return
}

func (uc *service) EditInformation(ctx echo.Context, payload dtos.InformationBodyRequest, id int64) (information dtos.InformationResponse, err error) {

	result, err := uc.model.GetInformationByID(ctx, id)
	if err != nil {
		return
	}

	_, err = uc.model.EditInformation(ctx, payload, id)
	if err != nil {
		return
	}

	information = dtos.InformationResponse{
		ID:       result.ID,
		Answer:   payload.Answer,
		Question: payload.Question,
	}

	return
}

func (uc *service) DeleteInformation(ctx echo.Context, id int64) (information dtos.InformationResponse, err error) {

	result, err := uc.model.GetInformationByID(ctx, id)
	if err != nil {
		return
	}

	_, err = uc.model.DeleteInformation(ctx, id)
	if err != nil {
		return
	}

	information = dtos.InformationResponse{
		ID:       result.ID,
		Answer:   result.Answer,
		Question: result.Question,
	}
	return
}

func (uc *service) GetInformations(ctx echo.Context) (informations []dtos.InformationResponse, err error) {
	result, err := uc.model.GetInformations(ctx)
	if err != nil {
		return
	}

	err = copier.Copy(&informations, result)
	return
}

func (uc *service) GetInformationByID(ctx echo.Context, id int64) (information dtos.InformationResponse, err error) {
	result, err := uc.model.GetInformationByID(ctx, id)
	if err != nil {
		return
	}

	err = copier.Copy(&information, result)
	return
}

func (uc *service) GetChatResponse(eCtx echo.Context, question string) (answer string, err error) {
    ctx := eCtx.Request().Context()
    informations, err := uc.model.GetInformations(eCtx)
    if err != nil {
        return "", err
    }
    
    if len(informations) == 0 {
        return "tidak ada data yang relevan di database kami", nil
    }

    model := uc.genai.GenerativeModel("gemini-pro")
    model.SetTemperature(0)
    
    parts := []genai.Part{
        genai.Text("Kamu adalah asisten yang membantu menjawab pertanyaan berdasarkan data yang tersedia. " +
            "Jika pertanyaan tidak relevan dengan data yang tersedia, jawab dengan: " +
            "'tidak ada data yang relevan di database kami'. " +
            "Berikut adalah data yang tersedia:\n\n"),
    }
    
    parts = append(parts, helpers.PopulateParts(informations, question)...)
    
    res, err := model.GenerateContent(ctx, parts...)
    if err != nil {
        return "", err
    }
    
    answer = helpers.ExtractAnswer(res)
    if answer == "" {
        return "tidak ada data yang relevan di database kami", nil
    }
    
    return answer, nil
}