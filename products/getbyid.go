package products

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getbyid struct {
	repository *gorm.DB
}

func (g getbyid) GetProductByID(ctx context.Context, id uint) (Product, error) {
	product := new(Product)

	if err := g.repository.WithContext(ctx).Preload(clause.Associations).First(product, "id = ?", id).Error; err != nil {
		return Product{}, err
	}

	return *product, nil
}
