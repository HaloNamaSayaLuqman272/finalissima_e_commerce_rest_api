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

func (g getbyrecommendation) GetProductByRecommendation(ctx context.Context, pagination utils.Pagination, productRecommendationsRequest *ai.ProductRecommendationRequest) (ai.ProductRecommendationResponse, error) {
	products := []Product{}
	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate)
}
