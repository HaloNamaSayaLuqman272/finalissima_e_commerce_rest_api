package products

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type Service interface {
	CreateProduct(ctx context.Context, createProductRequest *ProductRequest) (Product, error)
	GetProductByID(ctx context.Context, id uint) (Product, error)
	GetProductByCategory(ctx context.Context, pagination utils.Pagination, categoryId uint) (utils.Pagination, error)
	GetAllProduct(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	UpdateProduct(ctx context.Context, updateProductRequest *ProductRequest, id uint) (Product, error)
	DeleteProduct(ctx context.Context, id uint) error
}

type service struct {
	create
	getbyid
	getbycategory
	getall
	update
	delete
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		create:        create{repository: repository},
		getbyid:       getbyid{repository: repository},
		getbycategory: getbycategory{repository: repository},
		getall:        getall{repository: repository},
		update:        update{repository: repository, getbyid: getbyid{repository: repository}},
		delete:        delete{repository: repository, getbyid: getbyid{repository: repository}},
	}
}
