package utils

import (
	"fmt"
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty" query:"limit"`
	Page       int    `json:"page,omitempty" query:"page"`
	Sort       string `json:"sort,omitempty" query:"sort"`
	Search     string `json:"search,omitempty" query:"search"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Keyword    string `json:"keyword"`
	Rows       any    `json:"rows"`
}

func (p *Pagination) GetPage() int {
	// "func" kita mendeklarasikan fungsi
	// "p" variabel penerima dengan menerima pointer type struct "Pagination"
	// dengan nama fungsi "GetPage" tanpa ada parameter penerima
	// "int" fungsi ini akan mengembalikan bentuk "integer"
	if p.Page == 0 {
		// "p.Page" jika hasil dari variabel penerima "p" dari field "Page"
		// "== 0" masih kosong
		p.Page = 1
		// maka tetap munculkan field "Page" = 1 sebagai set default
	}

	return p.Page
	// mengembalikan hasil dari "p.Page"
}

func (p *Pagination) GetLimit() int {
	// / "func" kita mendeklarasikan fungsi
	// "p" variabel penerima dengan menerima pointer type struct "Pagination"
	// dengan nama fungsi "GetLimit" tanpa ada parameter penerima
	// "int" fungsi ini akan mengembalikan bentuk "integer"
	if p.Limit == 0 {
		// "p.Limit" jika hasil dari variabel penerima "p" dari field "Limit"
		// "== 0" masih kosong
		p.Limit = 10
		// maka tetap munculkan field "Limit" = 10 sebagi set default
	}

	return p.Limit
	// mengembalikan hasil dari "p.Limit"
}

func (p *Pagination) GetOffset() int {
	// "func" kita mendeklarasikan fungsi
	// "p" variabel penerima dengan menerima pointer type struct "Pagination"
	// dengan nama fungsi "GetOfset" tanpa ada parameter penerima
	// "int" fungsi ini akan mengembalikan bentuk "integer"
	return (p.GetPage() - 1) * p.GetLimit()
	// "p.GetPage() - 1" kita menuliskan "- 1" karena rumus "offset" biasa dimulai dari "0"
	// maka artinya mengembalikan setiap 1 page yg berisi limit 10
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id asc"
	}

	return p.Sort
}

func (p *Pagination) GetSearch() string {
	return p.Search
}

func Paginate(value any, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	// "func" kita mendeklarasikan fungsi
	// dengan nama "Paginate" tanpa memberi variabel penerima
	// dengan nama parameter "value" tipe data type assertion "any",
	// "pagination" tipe data pointer dari "Pagination",
	// dan "db" dengan tipe data package "GORM" method "DB"
	// "func(db *gorm.DB)" ini adalah pola "closure" atau "high-order function"
	// fungsi yg mengembalikan fungsi yg lain
	// kita menggunakan fungsi seperti ini karena kita menginginkan "Paginate"
	// kedepannya akan dipakai/ disisipkan pada query GORM yg lain
	// hal ini pada GORM memiliki method "Scopes()" yg memungkinkan untuk melakukan itu
	// misal penggunaan ketika kita menginginkan untuk menampilkan "Product"
	// maka dalam satu page akan berisi 10 "Product"
	// code: db.Scopes(Paginate(&products, &pagination, db)).Find(&products)
	// ibarat kata "Scopes()" adalah resep yg siap dipakai kapan saja
	var totalRows int64
	// kita membuat variabel dengan nama "totalRows" tipe data "int64"
	db.Model(value).Count(&totalRows)
	// "db.Model(value)" kita memanggil method "Model" yg bertugas
	// memberi tahu GORM tabel/ struct mana yg menjadi acuan untuk query berikutnya
	// dari koneksi database "db" (GORM)
	// "value" adalah parameter "any" yg dikirim ke fungsi "Paginate" sebelumnya
	// "Count" adalah method GORM yg menjalankan query "SELECT COUNT (*)"
	// ke tabel yg sudah ditentukan lewat ".Model(...)" yg bertugas untuk
	// menghitung total jumlah baris yg ada
	// "&totalRows" adalah alamat/ pointer dari "totalRows"
	// arti semuanya adalah tentukan tabel yg jadi acuan berdasarkan tipe data
	// "value" yg diberikan (misal tabel "product"), lalu hitung berapa
	// total baris yg ada di tabel itu, dan simpan hasil hitungannya ke variabel
	// "totalRows"

	pagination.TotalRows = totalRows
	// "totalRows" yg terisi hasil nya tadi dimasukkan ke dalam variabel
	// "pagination" di field "TotalRows"
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	// pada bagian code ini, kita memberi perintah untuk menampilkan
	// jumlah keseluruhan "page" dari hasil pembagian "totalRows" dengan "GetLimit"
	// seandainya "totalRows" berisi 23 baris dan "GetLimit" nya adalah 10,
	// maka akan ada 3 "page" dan di "page" 1 dan 2 akan terdapat 10 baris
	// dan "page" ke tiga terdapat 3 baris
	// "math" adalah package standar Go untuk opersai matematika
	// ".Ceil" adalah method pembulatan ke atas
	// "int" adalah konsep "totalPages" haruslah integer bilangan bulat
	// "totalPages" adalah variabel tempat penyimpanan hasil akhirnya
	pagination.TotalPages = totalPages
	// "totalPages" yg terisi hasil nya tadi dimasukkan ke dalam variabel
	// "pagination" di field "TotalPages"

	return func(db *gorm.DB) *gorm.DB {
		// mengembalikan fungsi "closure" dan tidak langsung mengembalikan paginate
		allowedColumns := map[string]bool{
			"name":        true,
			"description": true,
			"category":    true,
			"company":     true,
		}
		// membuat kolom yg akan query SQL terima proses di database
		// kolom yg akan diterima adalah kolom yg bernilai string
		// dan menyimpan daftar kolom yg diterima tersebut pada variabel
		// "allowedColumns"
		column := pagination.Keyword
		// menyimpan field "Keyword" dari struct "paginaton" di variabel "column"
		if !allowedColumns[column] {
			// seandainya kolom yg dicari tidak sesuai dengan kolom-kolom yg diterima
			// dalam variabel "allowedColumns"
			column = "name"
			// maka default kembalikan nilai kolom "name"
		}
		whereClause := fmt.Sprintf("%v LIKE ?", pagination.Keyword)
		// "fmt" adalah package standar Go
		// "Sprintf" merupakan salah satu method dari package "fmt" yg bertugas
		// membentuk string baru berdasarkan template tertentu, lalu mengembalikan
		// nilai stringg
		// "%v" adalah placeholder generic yg berfungsi untuk menyisipkan nilai apa saja
		// ke dalam bentuk default nya, yg nantinya akan digantikan argumen yg dikirm
		// "pagination.Keyword"
		// "LIKE" adalah operator SQL untuk pencocokan pola string
		// "?" merupakan placeholder untuk prepared statement, akan diisi nilai sebenarnnya
		// lewat parameter terpisah
		// "pagination.Keyword" mengakses field "Keyword" dari struct "pagination"
		// untuk mengisi bagian "%v"
		// "whereClause" nama variabel baru untuk menyimpan hasil string gabungan
		// (misalnya "product LIKE ?")
		// arti code ini bentuk sebuah klausa SQL "WHERE" secara dinamis, dengan
		// menyisipkan nama kolom yg ingin dicari (dari pagination.Keyword) ke dalam
		// template '<nama_kolom> LIKE ?'. Dan simpan hasil string ini ke variabel
		// "whereClause" untuk dipakai nanti di query "Where()"

		return db.Offset(pagination.GetOffset()).
			Limit(pagination.GetLimit()).
			Order(pagination.GetSort()).
			Where(whereClause, pagination.GetSearch()+"%")
		// arti dari code ini ambil koneksi database, terapkan aturan:
		// lewati skian baris (sesuai halaman), batasi jumlah baris yg
		// diambil (sesuai limit per halaman), urutkan hasilnya (sesuai kolom
		// dan arah yg ditentukan), lalu filter berdasarkan kondisi pencarian yg
		// sudah disiapkan. Kembalikan hasil query yg sudah lengkap
		// dengan semua aturan ini
	}
}

func PaginateByProductCategory(value any, pagination *Pagination, categoryId uint, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	// "func" kita mendeklarasikan sebuah fungsi baru yg bernama "PaginateByProductCategory"
	//  dengan nama parameter "value" tipe data type assertion "any",
	// "pagination" tipe data pointer dari "Pagination",
	// "categoryId" parameter penerima dengan tipe data uint
	// dan "db" dengan tipe data package "GORM" method "DB"
	// "func(db *gorm.DB)" ini adalah pola "closure" atau "high-order function"
	// fungsi yg mengembalikan fungsi yg lain
	// kita menggunakan fungsi seperti ini karena kita menginginkan "PaginateByProductCategory"
	// kedepannya akan dipakai/ disisipkan pada query GORM yg lain
	var totalRows int64
	whereCategoryClause := "category_id = ?"

	db.Model(value).Where(whereCategoryClause, categoryId).Count(&totalRows)
	// ".Where(...)" menambahkan kondisi filter ke query
	// "whereCategoryClause" string kondisi SQL
	// "categoryId" nilai yg akan mengisi "?" dari "category_id = ?"
	// ".Count" menjalankan "SELECT COUNT(*)" tapi mengikuti kondisi "WHERE"
	// yg sudah ditambahkan sebelumnya
	// "&totalRows" adalah pointer variabel "totalRows" dan akan menjadi tempat
	// menyimpan hasil hitungan

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		whereClause := fmt.Sprintf("%v LIKE ?", pagination.Keyword)

		return db.Offset(pagination.GetOffset()).
			Limit(pagination.GetLimit()).
			Order(pagination.GetSort()).
			Where(whereClause, pagination.GetSearch()+"%").
			Where(whereCategoryClause, categoryId)
		// kita menambahkan secara spesifik ".Where(whereCategoryClause, ...)"
		// karena pada kondisi ini, kita ingin Go memproses dan menampilkan
		// pagination secara khusus
		// jika yg sebelumnya misalnya ingin menampilkan daftar baris "Product"
		// dan akan menampilkan semua "Product" yg ada,
		// di bagian ini akan diproses lebih spesifik "Category Product",
		// seperti misalnya kategori produk "Halal Food" maka yg akan ditampilkan
		// hanyalah produk dari kategori "Halal Food"
	}
}
