package categories

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type Service interface {
	CreateCategory(ctx context.Context, createCategoryRequest *CategoryRequest) (Category, error)
	GetByID(ctx context.Context, id uint) (Category, error)
	GetAllCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error)
	UpdateCategory(ctx context.Context, updateCategoryRequest *CategoryRequest, id uint) (Category, error)
	DeleteCategoryByID(ctx context.Context, id uint) error
	// kita bungkus semua fungsi yg ada pada fitur "Category" beserta parameter input
	// dan variabel kembalian nya, kemudian kita beri nama type interface ini
	// dengan "Service"
	// kita membuat ini untuk kebutuhan "mockery" yg akan digunakan sebagai
	// unit testing code
}

type service struct {
	create
	get
	getall
	update
	delete
	// tidak lupa kita bungkus juga semua field yg ada pada fitur "Category"
	// dan kita namakan type struct ini dengan nama "service"
	// kita membutuhkan sesuatu yg menyatukan kelima field dengan tujuan
	// agar bisa dipanggil dengan lebih praktis, alih-alih membuat pola
	// struktur yg berulang-ulang
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB) Service {
	return service{
		create: create{repository: repository},
		get:    get{repository: repository},
		getall: getall{repository: repository},
		update: update{repository: repository, get: get{repository: repository}},
		delete: delete{repository: repository, get: get{repository: repository}},
	}
	// ini adalah "constructor" tujuannya membuat dan menyiapkan instance "service"
	// yg sudah lengkap dan siap pakai, tanpa pemanggil perlu tahu detail
	// cara menyusunnya secara detail
}
