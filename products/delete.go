package products

import (
	"context"

	"gorm.io/gorm"
)

type delete struct {
	repository *gorm.DB
	getbyid
}

func (d delete) DeleteProductByID(ctx context.Context, id uint) error {
	product, err := d.GetProductByID(ctx, id)
	if err != nil {
		return err
	}

	if err := d.repository.WithContext(ctx).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
