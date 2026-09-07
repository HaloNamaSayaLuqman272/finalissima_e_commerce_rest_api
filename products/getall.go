package products

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllProduct(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	products := []Product{}
}
