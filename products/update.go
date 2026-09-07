package products

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"

	"gorm.io/gorm"
)

type update struct {
	repository *gorm.DB
	getbyid
}

func (u update) UpdateProduct(ctx context.Context, updateProductRequest *ProductRequest, id uint) (Product, error) {
	product := models.Product{
		NameProduct: updateProductRequest.NameProduct,
		CategoryID:  updateProductRequest.CategoryID,
		Description: updateProductRequest.Description,
		Company:     updateProductRequest.Company,
		Barcode:     updateProductRequest.Barcode,
		Weight:      updateProductRequest.Weight,
		ExpiredDate: updateProductRequest.ExpiredDate,
		Price:       updateProductRequest.Price,
		Stock:       updateProductRequest.Stock,
		ImageLink:   updateProductRequest.ImageLink,
	}

	result := u.repository.WithContext(ctx).Where("id = ?", id).Updates(&product)
	if err := result.Error; err != nil {
		return Product{}, err
	}

	record, err := u.getbyid.GetProductByID(ctx, id)
	if err != nil {
		return Product{}, err
	}

	return record, nil
}
