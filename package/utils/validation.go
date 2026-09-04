package utils

import "finalissima_e_commerce_rest_api/package/constant"

func ValidateFile(extension string) bool {
	// "func" kita mendeklarasikan fungsi
	// selanjutnya menamai fungsi dengan "ValidateFile"
	// dan kita tidak meembutuhkan variabel penerima di depannya,
	// berarti ini adalah fungsi biasa, bukan method,
	// yg bisa dipanggil langsung lewat nama package "utils",
	// exp: utils.ValidateFile("file.txt")
	// "extension string" adalah parameter input fungsi ini,
	// dan bertipe data string
	// kita membutuhkan parameter yg diterima agar file diproses
	// dan dikembalikan dengan mudah diidentifikasi
	// "bool" adalah tipe data yg dikembalikan fungsi ini,
	// yaitu boolean, artinya fungsi ini akan mengembalikan nilai true atau false
	// kita membutuhkan validasi file yg diterima dari user
	// rencanaya adalah untuk gambar product dan gambar user profile
	return constant.ALLOWED_EXTENSIONS[extension]
	// mengembalikan variabel "ALLOWED_EXTENSIONS" dari package "constant"
	// dan hasil dari parameter "extension" apakah nilainya "true" atau "false"
}
