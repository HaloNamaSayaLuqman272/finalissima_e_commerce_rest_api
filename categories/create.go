package categories

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"

	"gorm.io/gorm"
)

type create struct {
	// kita perlu membuat type struct "create"
	// yg akan kita gunakan sebagai variabel penerima pada function "CreateCategory"
	repository *gorm.DB
	// kita membutuhkan "repository" yg terhubung dengan gorm.DB
}

func (c create) CreateCategory(ctx context.Context, createCategoryRequest *CategoryRequest) (Category, error) {
	category := models.Category{
		NameCategory: createCategoryRequest.NameCategory,
		Description:  createCategoryRequest.Description,
	}
	// kita membuat variabel "category" yg akan menampung data dari request body

	result := c.repository.WithContext(ctx).Create(&category)
	// kita membuat variabel "result" yg akan menampung hasil dari proses create data category
	record := new(Category)
	// selanjutnya kita membuat variabel "record" yg akan menampung data category yg telah berhasil dibuat

	if err := result.Error; err != nil {
		// jika variabel "result" error maka akan menampilkan error
		return Category{}, err
	}

	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		// "result" adalah variabel *gorm.DB yg menampung hasil dari proses create data category
		// "WithContext(ctx)" untuk mengikat "ctx" agar tunduk pada lifecycle request, bisa timeout atau dibatalkan
		// "Last(record)" method GORM untuk mengambil data terakhir yg baru saja dibuat
		// kemudan hasilnya dimasukkan ke dalam variabel pointer "record"
		// "Error" mengambil field "Error" dari hasil operasi "Last",
		// ini akan terisi otomatis apabila terjadi masalah saat query dijalankan
		return Category{}, err
		// mengembalikan "Category{}" kosongan dan menampilkan error jika terjadi masalah saat query dijalankan
	}

	return *record, nil
	// jika berhasil maka kembalikan nilai hasil inputan "record"
	// dan  "nil" bersih dari "error"
}
