package products

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"

	"gorm.io/gorm"
)

type create struct {
	repository *gorm.DB
}

func (c create) CreateProduct(ctx context.Context, createProductRequest *ProductRequest) (Product, error) {
	product := models.Product{
		NameProduct: createProductRequest.NameProduct,
		CategoryID:  createProductRequest.CategoryID,
		Description: createProductRequest.Description,
		Company:     createProductRequest.Company,
		Barcode:     createProductRequest.Barcode,
		Weight:      createProductRequest.Weight,
		ExpiredDate: createProductRequest.ExpiredDate,
		Price:       createProductRequest.Price,
		ImageLink:   createProductRequest.ImageLink,
	}

	result := c.repository.WithContext(ctx).Create(&product)
	record := new(Product)
	if err := result.Error; err != nil {
		return Product{}, err
	}

	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		return Product{}, err
	}

	return *record, nil
}
