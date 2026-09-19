package purchases

import (
	"context"

	"gorm.io/gorm"
)

type getbypurchaseid struct {
	repository *gorm.DB
}

func (g getbypurchaseid) GetByPurchaseID(ctx context.Context, id uint) (Purchase, error) {
	purchase := new(Purchase)
	if err := g.repository.WithContext(ctx).Preload("User").Preload("Product").Preload("Product.Category").First(purchase, "id = ?", id).Error; err != nil {
		return Purchase{}, err
	}

	return *purchase, nil
}
