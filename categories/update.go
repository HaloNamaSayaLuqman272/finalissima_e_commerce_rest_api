package categories

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	get
}

// dalam type struct "update" terdapat field "repository" dengan tipe data
// pointer ke GORM
// dan ada field "get" yg akan kita gunakan untuk mengambil data "id" mana
// yg akan di-update

func (u update) UpadateCategory(ctx context.Context, categoryRequest *CategoryRequest, id uint) (Category, error) {
	category := models.Category{
		NameCategory: categoryRequest.NameCategory,
		Description:  categoryRequest.Description,
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&category)
	if err := result.Error; err != nil {
		return Category{}, err
	}

	record, err := u.get.GetByID(ctx, id)
	// panggil method "GetByID" yg dimiliki field "get" di dalam struct "u",
	// kirim "ctx" dan "id" sebagai parameter. Simpan hasil datanya ke variabel
	// "record" dan apabila ada error simpan di variabel "err"
	if err != nil {
		return Category{}, err
	}

	return record, nil
}
