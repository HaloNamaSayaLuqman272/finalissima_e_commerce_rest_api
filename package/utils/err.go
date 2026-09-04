package utils

import "strings"

// kita membuat package utils sebagai tempat menyimpan fungsi-fungsi
// yg bersifat umum dan generik, yg sering dipakai di banyak tempat,
// namun tidak spesifik untuk satu domain tertentu
func GetValidationErrorMessage(message string) []string {
	// "func" kita mendeklarasikan fungsi
	// "GetValidationErrorMessage" adalah nama fungsinya
	// dan kita tidak meembutuhkan variabel penerima di depannya,
	// berarti ini adalah fungsi biasa, bukan method,
	// yg bisa dipanggil langsung lewat nama package "utils",
	// exp: utils.GetValidationErrorMessage("error message")
	// "message string" adalah parameter input fungsi ini,
	// dan bertipe data string
	// "[]string" adalah tipe data yg dikembalikan fungsi ini,
	// yaitu slice of string, artinya fungsi ini akan memecah satu pesan yg panjang
	// menjadi beberapa pesan yg lebih pendek, dan mengembalikannya
	separator := ","
	// saat hasil data dipecah menjadi "[]string" normalnya tidak ada elemen pemisah
	// murni hanya ditandai dengan spasi kosong
	// maka disini kita perlu menambah pemisah di akhir string
	// dan kita namakan ini dengan variabel "separator"
	// dengan isi ",", artinya setiap pesan akan dipisahkan dengan tanda koma
	errors := strings.Split(message, separator)
	// "errors" adalah variabel yg menampung hasil pemecahan string
	// "strings.Split" adalah fungsi bawaan dari package "strings"
	// dan "Split" adalah salah satu method nya,
	// yg berfungsi untuk memecah satu string menjadi beberapa bagian
	// "message" adalah string yg akan dipecah
	// "separator" adalah string yg menjadi pemisah antara satu bagian dengan bagian lainnya
	// hasilnya akan dikembalikan dalam bentuk slice of string, dan disimpan di variabel "errors"
	return errors[:len(errors)-1]
	// mengembalikan variabel "errors" yg berisi hasil pemecahan string
	// "[:len(errors)-1]" adalah slicing, artinya kita mengambil semua elemen dari slice "errors"
	// ini adalah proses slicing
	// "len(errors)" adalah fungsi bawaan dari package "len"
	// yg berfungsi untuk menghitung jumlah elemen dalam slice "errors"
	// dan kita mengurangi 1 dari hasilnya, karena kita ingin menghilangkan elemen terakhir
	// yg biasanya kosong, karena pemisah terakhir tidak diikuti oleh pesan
}
