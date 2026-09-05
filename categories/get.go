package categories

import (
	"context"

	"gorm.io/gorm"
)

type get struct {
	repository *gorm.DB
}

func (g get) GetByID(ctx context.Context, id uint) (Category, error) {
	// kita mendeklarasikan sebuah fungsi dan memberi nama "GetByID"
	// memiliki variabel penerima "g" bertipe data "get"
	// fungsi "GetByID" memiliki variabel input "ctx" bertipe data "context.Context"
	// dan "id" bertipe data "uint"
	// rencananya akan mengembalikan hasil data "Category" dan "error"apabila
	// bermasalah
	category := new(Category)
	// "new" adalah fungsi bawaan Go yg mengalokasikan memori kosong (zero value)
	// untuk suatu tipe, lalu mengembalikan pointer ke memori itu
	// "*Category" yg kosong akan diisi nanti lewat proses query database
	// kemudian hasil dari proses fungsi bawaan ini akan disimpan di
	// variabel "category"

	if err := g.repository.WithContext(ctx).First(category, "id = ?", id).Error; err != nil {
		// kita membuat alur mendapatkan "category" berdasarkan "id"
		// "g.repostory.WithContext(ctx)" code ini untuk mendukung manajemen
		// lifecycle request "get"
		return Category{}, err
		// kondisi jika gagal akan mengembalikan struct "Category" kosongan
		// dan penyebab error
	}

	return *category, nil
	// kondisi jika berhasil akan mengembalikan pointer ke variabel "category"
	// dan hasilnya "nil"
}
