package categories

import (
	"context"

	"gorm.io/gorm"
)

type delete struct {
	repository *gorm.DB
	get
	// kita membutuhkan field "get" untuk menentukan "category_id" mana
	// yg akan dihapus
}

func (d delete) DeleteCategoryByID(ctx context.Context, id uint) error {
	category, err := d.GetByID(ctx, id)
	// bagian ini adalah bagian memasukkan "category_id" yg akan dihapus
	if err != nil {
		return err
	}

	if err := d.repository.WithContext(ctx).Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
