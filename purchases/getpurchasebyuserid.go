package purchases

import (
	"context"

	"gorm.io/gorm"
)

type getpurchasebyuserid struct {
	repository *gorm.DB
}

func (g getpurchasebyuserid) GetPurchaseByUserID(ctx context.Context, userId uint) ([]Purchase, error) {
	purchases := []Purchase{}
	if err := g.repository.WithContext(ctx).Preload("User").Preload("Product").Preload("Product.Category").Where("user_id = ?", userId).Find(&purchases).Error; err != nil {
		return nil, err
	}

	return purchases, nil
}
