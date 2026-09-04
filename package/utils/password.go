package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) ([]byte, error) {
	// kita mendeklarasikan sebuah fungsi "func"
	// dengan nama "GeneratePassword" dan parameter input "password"
	// tipe data "string"
	// "([]byte, error)" mengembalikan slice of byte dan "error" apabila ada masalah
	return bcrypt.GenerateFromPassword(
		// kita langsung melakukan pengembalian fungsi lain karena
		// "GeneratePassword" dan "GenerateFromPassword" sama-sama mengembalikan
		// dua nilai "[]byte dan error"
		// ini diperbolehkan dan lebih simpel, alih-alih memakai cara manual yg panjang
		[]byte(password), bcrypt.DefaultCost,
		// fungsi "GenerateFromPassword" dengan parameter input
		// "[]byte" sebagai konversi "password" yg menjadi pengisi parameter
		// dan "bcrypt.DefaultCost" adalah konstanta standar untuk proses
		// hashing data
		// "[]byte(password)" adalah apa yg akan di-hash (password)
		// dan "DefaultCost" adalah seberapa berat cara hash yg akan dilakukan
		// arti code panggil fungsi "GenerateFromPassword" dari package "bcrypt",
		// kirim "password" (yg sudah dikonversi []byte) sebagai bahan yg akan di-hash
		// dan gunakan hashing standar (DefaultCost) dari package tersebut
	)
}

func ComparePassword(hashed, password string) error {
	// kita akan mendeklarasikan sebuah fungsi untuk melakukan komparasi
	// antara "hashed" password yg berhasil di-hashed
	// dan "password" password yg belum di-hashed
	// kita namakan "ComparePassword" dengan parameter input "hashed" dan
	// "password" sama-sama bertipe data string
	// fungsi ini akan mengembalikan error
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed), []byte(password),
		// arti code panggil fungsi "CompareHashAndPassword" dari package "bcrypt",
		// kirim "hashed" (password yg sudah di-hashed dan sudah dikonversi []byte)
		// dan kirim "password" (yg sudah dikonversi []byte)
	)
}
