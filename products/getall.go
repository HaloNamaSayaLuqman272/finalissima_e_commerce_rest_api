package products

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type getall struct {
	repository *gorm.DB
}

func (g getall) GetAllProduct(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	products := []Product{}

	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&products, &pagination, g.repository)).Preload(clause.Associations).Find(&products).Error; err != nil {
		// "g.repository.WithContext(ctx)" untuk menghubungkan ke database dan
		// ".WithContext(ctx)" menjaga proses lifecycle aplikasi Go, jika
		// dari user melakukan cancel atau sambungan database ada masalah
		// maka akan proses ini akan dihentikan
		// ".Scopes(...)" kita ingin menggunakan "Pagination" dari package
		// "utils" untuk menampilkan semua "products" yg ada di database
		// dan tampilkan halaman per halaman
		// ".Preload(clause.Associations)" adalah cara memuat semua data relasi/
		// asosiasi sekaligus secara otomatis
		return utils.Pagination{}, err
		// jika terdapat masalah atau error saat proses pengambilan maka kembalikan
		// "Pagination" kosongan dan error
	}

	pagination.Rows = products

	return pagination, nil
}
