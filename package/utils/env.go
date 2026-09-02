package utils

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// di package "utils" ini kita membuat "env"
// untuk menjelaskan kepada Go tentang alur memproses "env"
// dan menyimpan konfigurasi dan data sensitif
func GetConfigurance(key string) string {
	// "func" kita mendeklarasikan adanya fungsi
	// dan kita namakan fungsi "GetConfigurance"
	// "key string" adalah parameter penerima yg bernama "key" dengan tipe data "string"
	// "string" mengembalikan tipe data string
	isDevelop := os.Getenv("APP_MODE") != "production"
	// "os" adalah package standar Go untuk berinteraksi dengan sistem operasi
	// "Getenv" adalah salah satu method dari package "os"
	// yg bertugas membaca nilai dari variabel environment berdasarkan nama "key" nya
	// "APP_MODE" adalah nama variabel environment yg akan diambil nilainya
	// "production" adalah string yg dibandingkan
	// analoginya apabila "APP_MODE" nilai nya tidak sama dengan "production"
	// maka hasil perbandingan nya "true"
	// "isDevelop" adalah variabel baru yg secara otomatis bertipe "bool"
	// akan bernilai "true" jika aplikasi berjalan di selain mode "production"
	// seperti mode "development" atau "testing"
	// dan akan bernilai "false" apabila berjalan di mode "production"
	// arti keseluruhan nya adalah baca nilai environment variabel "APP_MODE",
	// cek apakah tidak bernilai "production".
	// simpan hasil pengecekan (true/ false) ke variabel "isDevelop"
	// yg menandakan apakah aplikasi sedang berjalan dalam mode "production" atau tidak

	if isDevelop {
		// disini kita membuat percabangan untuk mencegah kode mencoba mencari
		// file yg memang tidak akan pernah ada di "production"
		// ingat ".env" akan kita larang untuk dideploy melalui ".gitignore"
		// sehingga kita perlu sesuatu yg bisa membaca/ deploy ".env" saat masih di local/ development
		viper.AddConfigPath(".")
		// "viper" adalah package standar Go yg berfungsi untuk membaca isi file ".env"
		// menyuntikkan nilainya agar bisa diakses seperti variabel biasa environment
		// "AddConfigPath" adalah method dari package "viper" yg tugasnya
		// memberitahu "viper" folder mana yg harus dicari
		// "." artinya folder root project
		viper.SetConfigFile(".env")
		// "SetConfigFile" adalah method dari package "viper" yg berfungsi
		// meberitahu package "viper" file konfigurasi mana yg harus dicari
		// ".env" adalah file konfigurasi yg dicari

		if err := viper.ReadInConfig(); err != nil {
			// "ReadInConfig" adalah salah satu method "viper" yg bertugas
			// benar-benar membaca isi file yg ditunjuk lokasinya lewat
			// method "AddConfigPath" dan "SetInConfig"
			// isi file akan dibaca dan di-parse
			// proses ini akan mengembalikan "error"
			// "err" adalah variabel yg akan menjadi tempat menyimpan "error"
			// "err != nil" adalah kondisi jika "err" tidak bernilai "nil"
			log.Fatalf("error when reading configuration file: %v\n", err)
			// "log" adalah package standar Go untuk keperluan mencetak pesan di terminal
			// "Fatalf" adalah method dari package "log" yg memiliki tugas
			// mencetak pesan, dan setelah pesan tercetak fungsi ini akan menghentikan
			// program karena ada masalah, bukan berhenti secara normal
			// "err" mengambil variabel "err" untuk mengisi placeholder "%v"
		}

		return viper.GetString(key)
		// "GetString" adalah salah satu method package "viper" yg bertugas
		// mengambil satu nilai konfigurasi berdasarkan key nya dari data
		// yg sudah berhasil dibaca sebelumnya lewat "ReadInConfig" lalu
		// mengembalikannya ke bentuk tipe string
		// exp: dbHost := viper.GetString("DB_HOST") -> hasilnya: "localhost"
	}

	return os.Getenv(key)
	// mengembalikan nilai dari environment parameter "key"
}
