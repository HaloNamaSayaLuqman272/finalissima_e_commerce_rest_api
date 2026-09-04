package utils

import "gorm.io/gorm"

func CurrentUser(userID uint) func(db *gorm.DB) *gorm.DB {
	// kita mendeklarasikan sebuah fungsi biasa dengan nama "CurrentUser"
	// parameter input bernama "UserID" tipe data "uint"
	// kemudian fungsi "CurrentUser" akan mengembalikan fungsi lain
	return func(db *gorm.DB) *gorm.DB {
		// langsung mengembalikan fungsi lain
		return db.Where("user_id = ?", userID)
		// mengembalikan database dari "UserID"
	}
	// arti code buat sebuah fungsi filteryg bisa dipasang di query GORM manapun
	// yg secara otomatis membatasi hasil query hanya untuk baris-baris yg "user_id"
	// cocok dengan ID user yg diberikan
}

// tujuan kita perlu membuat ini
// 1. Akan menjadi endpoint data milik saya sendiri, contoh code
// db.Scopes(utils.CurrentUser(userID)).Find(&orders)
// artinya tampilkan semua order milik user tertentu
// setara "SELECT * FROM orders WHERE user_id = <userID>"
// 2. Bisa digabung dengan scope lain, contoh code
// db.Scopes(utils.CurrentUser(userID), utils.Paginate(&orders, &pagination, db),).Find(&order)
// artinya tampilkan daftar order milik user dengan ID tertentu, dan tampilkan saja
// sesuai dengan halaman yg diminta (bukan menampilkan  sekaligus semua order yg ada)
// "pagination" memiliki kata kunci sehingga menjadi urut dan ter-filter
// simpan hasilnya ke "orders"
// setara "SELECT * FROM orders
// 	WHERE user_id = <userID>
// 		AND <kondisi search dari pagination, kalau ada>
// 	ORDER BY <sort dari pagination>
// 	LIMIT <limit dari pagination>
// 	OFFSET <offset dari pagination>"
// 3. Untuk kebutuhan keamanan, memastikan user A tidak bisa memanipulasi
// data user B, begitu pula sebaliknya dengan cara menyisipkan filter "user_id"
// di query
