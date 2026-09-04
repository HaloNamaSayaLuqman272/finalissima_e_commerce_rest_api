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
	//

}
