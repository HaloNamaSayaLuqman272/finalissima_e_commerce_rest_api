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
		// bagian ini adalah bagian daftar data yg perlu user isi
		// dan struktur model data ini akan dibungkus di variabel "product"
	}

	result := c.repository.WithContext(ctx).Create(&product)
	// code ini berisi tentang proses pengisian data variabel "product"
	// "c.repository" adalah titik masuk ke database, kita menggunakan
	// koneksi database yg yg sudah disiapkan struct "create"
	// alur koneksi dan pengisian data "create product" kita bungkus
	// menggunakan variabel dan bernama "result"
	if err := result.Error; err != nil {
		return Product{}, err
		// seandainya variabel "result" terdapat masalah, maka akan mengembalikan
		// struct "Product{}" kosongan dan menampilkan masalah error
	}

	record := new(Product)
	// fungsi "new" yg berparameter input terbaru "Product" ini akan dibungkus
	// dalam variabel "record"
	// ingat, variabel "record" ini berisi baris-baris data "Product" yg
	// aslinya masih kosong dan akan terisi pada pemanggilan "record"
	// di ".Last(record)"
	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		// code ini adalah proses pengambilan data terakhir "create product"
		// dan dimasukkan ke dalam variabel "record"
		return Product{}, err
		// seandainya proses mengambil data terakhir "record" terjadi masalah
		// maka akan mengembalikan struct "Product{}" kosongan dan menampilkan
		// error
	}

	return *record, nil
	// keadaan berhasil akan mengembalikan pointer ke variabel "record" dan
	// "nil" bersih dari error
}
