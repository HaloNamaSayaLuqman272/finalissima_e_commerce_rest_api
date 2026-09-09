package ai

const BASE_URL = "https://openrouter.ai/api/v1"

type Service interface {
	GetProductRecommendation()
}
