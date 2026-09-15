package ai

import (
	"encoding/json"
	"finalissima_e_commerce_rest_api/package/clients"
	"finalissima_e_commerce_rest_api/package/constant"
	"finalissima_e_commerce_rest_api/package/utils"
	"fmt"
	"net/http"
)

const BASE_URL = "https://openrouter.ai/api/v1"

type Service interface {
	GetProductRecommendation(req ProductRecommendationRequest) (PromptResponse, error)
}

type service struct {
	client clients.HTTPClient
}

func InitService() Service {
	return &service{
		client: clients.InitHTTPClient(BASE_URL, 40, utils.GetConfigurance(constant.AI_API_KEY)),
	}
}

func (r *service) GetProductRecommendation(req ProductRecommendationRequest) (PromptResponse, error) {
	var response PromptResponse
	var model string = utils.GetConfigurance(constant.AI_MODEL)
	var userPrompt string = fmt.Sprintf("Suggest TOP %v PRODUCT RECOMMENDATIONS about %v", req.Quantity, req.Topic)

	payload := map[string]any{
		"model": model,
		"message": []Message{
			{
				Role:    "system",
				Content: SYSTEM_PROMPT,
			},
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
	}

	res, err := r.client.SendJSON("/chat/completions", http.MethodPost, payload)
	if err != nil {
		return PromptResponse{}, err
	}

	if err := json.Unmarshal([]byte(res), &response); err != nil {
		return PromptResponse{}, err
	}

	return response, nil
}
