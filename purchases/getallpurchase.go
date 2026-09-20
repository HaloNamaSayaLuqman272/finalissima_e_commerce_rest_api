package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type getallpurchase struct {
	repository *gorm.DB
}

func (g getallpurchase) GetAllPurchase(ctx context.Context, pagination utils.Pagination) ([]Purchase, error) {
	purchases := []Purchase{}
	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&purchases, &pagination, g.repository)).Preload("User").Preload("Product").Preload("Product.Category").Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}
