package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type getallpurchase struct {
	repository *gorm.DB
}

func (g getallpurchase) GetAllPurchase(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	purchases := []Purchase{}
}
