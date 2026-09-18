package products

import (
	"context"
	"encoding/json"
	"errors"
	"finalissima_e_commerce_rest_api/package/ai"
	"finalissima_e_commerce_rest_api/package/utils"
	"log"
	"math"

	"gorm.io/gorm"
)

type getbyrecommendation struct {
	repository *gorm.DB
}

func (g getbyrecommendation) GetProductByRecommendation(ctx context.Context, pagination utils.Pagination, promptRequest ai.ProductRecommendationRequest) (utils.Pagination, error) {
	// kirim semua product yg ada di database ke AI
	var allProducts []Product
	if err := g.repository.WithContext(ctx).Find(&allProducts).Error; err != nil {
		return utils.Pagination{}, err
	}
	if len(allProducts) == 0 {
		return utils.Pagination{}, errors.New("no products available")
	}

	// ubah ke struct ringan milik AI
	aiProducts := make([]ai.Product, len(allProducts))
	for i, p := range allProducts {
		aiProducts[i] = ai.Product{
			ID:          p.ID,
			NameProduct: p.NameProduct,
			Description: p.Description,
			CategoryID:  p.CategoryID,
			Price:       p.Price,
		}
	}
	promptRequest.Products = aiProducts

	// minta AI memilih dari daftar itu
	aiService := ai.InitService()
	res, err := aiService.GetProductRecommendation(promptRequest)
	if err != nil {
		return utils.Pagination{}, err
	}

	resBody := res.Choices[0].Message.Content
	log.Println("DEBUG AI raw response:", resBody)
	var selected []struct {
		ID uint `json:"id"`
	}
	if err := json.Unmarshal([]byte(resBody), &selected); err != nil {
		return utils.Pagination{}, err
	}
	if len(selected) == 0 {
		return utils.Pagination{}, errors.New("no products selected by AI")
	}

	ids := make([]uint, len(selected))
	for i, s := range selected {
		ids[i] = s.ID
	}

	// ambil data lengkap product sesuai pilihan AI
	var totalRows int64
	if err := g.repository.Model(&Product{}).Where("id IN ?", ids).Count(&totalRows).Error; err != nil {
		return utils.Pagination{}, err
	}
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))

	var recommendedProducts []Product
	if err := g.repository.WithContext(ctx).Where("id IN ?", ids).Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Preload("Category").Find(&recommendedProducts).Error; err != nil {
		return utils.Pagination{}, err
	}

	pagination.Rows = recommendedProducts
	return pagination, nil
}
