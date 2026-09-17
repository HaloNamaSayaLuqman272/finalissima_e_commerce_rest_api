package products

import (
	"context"
	"finalissima_e_commerce_rest_api/package/ai"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type getbyrecommendation struct {
	repository *gorm.DB
}

func (g getbyrecommendation) GetProductByRecommendation(ctx context.Context, pagination utils.Pagination, promptRequest ai.ProductRecommendationRequest) (utils.Pagination, error)
