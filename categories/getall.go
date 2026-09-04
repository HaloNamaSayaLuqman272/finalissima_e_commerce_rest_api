package categories

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type getall struct {
	// kita membuat type struct bernama "getall" yg rencana nya akan digunakan
	// untuk menampilkan semua data kategori yg ada
	repository *gorm.DB
	create
	// dalam type struct "getall" kita membutuhkan "repository" yg terhubung dengan *gorm.DB
	// dan juga kita ingin type struct "getall" mendapatkan akses koneksi ke "create"
}

func (g getall) GetAllCategories(ctx context.Context, pagination utils.Pagination) (utils.Pagination, error) {
	// kita mendeklarasikan sebuah fungsi dengan nama "GetAllCategories"
	// memiliki variabel penerima "g" bertipe data "getall"
	// terdapat parameter input "ctx" tipe data dari "context.Context" untuk menjaga lifcycle
	// dan "pagination" tipe data dari "utils.Pagination" yg akan menampilkan
	// jumlah halaman dan per halaman akan menampilkan maksimal berapa baris data
	categories := []Category{}
	// "[]Category{}" adalah slice of Category{} yg rencana nya
	// akan menjadi tempat menampung banyak kategori, saat ini masih kosongan
	// dan "[]Category{}" akan dibungkus variabel "categories"

	if err := g.repository.WithContext(ctx).Scopes(utils.Paginate(&categories, &pagination, g.repository)).Find(&categories).Error; err != nil {
		// selanjutnya kita membuat alur "GetAllCategory", kita mengnginkan
		// Go agar menampilkan "semua kategori yg ada", menggunakan pengaturan
		// halaman dari "pagination", dan menambahkan "g.repository"
		// untuk menyambungkan ke database
		// ".Find(&categories)" kemudian kita menjalankan perintah "Find"
		// untuk mencari semua hasil "categories" dan ambil semua data nya
		// dan mengambil field "Error" dari proses operasi "Find"
		// apabila terjadi masalah saat query dijalankan
		// lalu kita bungkus ini ke dalam variabel "err";
		// namun bila variabel "err" tidak bernilai "nil" atau tidak bersih dari error
		return utils.Pagination{}, err
		//maka kembalikan "Pagination{}" kosongan dan hasil error
		// kita melakukan ini agar memastikan pemanggil fungsi tidak keliru
		// menganggap data "pagination" yg error itu valid dan "err" akan
		// memberitahu penyebab error nya agar bisa ditangani dengan tepat
	}

	pagination.Rows = categories
	// hasil dari "categories" akan disimpan pada field "Rows" dari struct "pagination"

	return pagination, nil
	// jika "GetAllCategories" berhasil maka akan mengembalikan isi "pagination"
	// dan "nil"
}
