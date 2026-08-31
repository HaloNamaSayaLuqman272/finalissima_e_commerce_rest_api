package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

// kita memulai project dengan membuat bentuk models yg diharapkan
// selanjutnya membuat type struct "User" berisi apa saja data yg diperlukan
// kita membutuhkan data apa saja yg perlu user isi untuk project kita ini
// karena rencananya akan ada banyak data yg diperlukan,
// maka kita bungkus dalam type struct untuk menampung nama data-data tersebut
// setidaknya;
// kita perlu nomer ID yg unik
// username tentunya
// email
// password
// phone_number untuk kontak sewaktu-waktu dibutuhkan lebih cepat jika terjadi sesuatu
// address untuk alamat jalan, nomer, dan nama komplek rumah jika ada
// province_id, sesuai dengan provinsi
// city_id, sesuai dengan kota kabupaten
// district_id, sesuai denagn kota kecamatan
// role -> default akan menjadi user
// created_at, menunjukkan waktu kapan registrasi sukses dibuat
// updated_at, menunjukkan waktu kapan update profile dilakukan
// deleted_at, menunjukkan waktu kapan user dihapus dari sistem
type User struct {
	ID uint `json:"id" form:"-" gorm:"primaryKey;autoIncrement"`
	// primaryKey -> field ini akan dijadikan sebagai kunci utama tabel SQL ||
	// autoIncrement -> perintah untuk database agar mengisi ID otomatis sesuai dengan angka urutan
	// tanda (-) -> abaikan field ini dan user tidak bisa mengisi ID sendiri
	// json -> menyertakan ID agar sistem mengembalikan respon pada frontend jika registrasi user berhasil
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	// kita akan melakukan validasi password di tempat terpisah
	PhoneNumber string `json:"phone_number" gorm:"uniqueIndex"`
	// kita juga akan melakukan validasi nomer telepon di tempat terpisah
	// not null -> data tidak boleh kosong oleh user
	// size:20 -> jumlah digit maksimal 20
	Address        string         `json:"address" gorm:"type:text"`
	ProvinceID     uint           `json:"province_id"`
	CityID         uint           `json:"city_id"`
	DistrictID     uint           `json:"district_id"`
	ProfilePicture string         `json:"profile_picture"`
	Role           string         `json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	// alih-alih menggunakan time.Time, menggunakan gorm.DeletedAt
	// agar sistem tetap menampilkan data yg di-delete
	// dan menampilkan NULL pada isi database
	// ini adalah soft delete
	// artinya jika sewaktu-waktu ingin melihat atau mengembalikan data, maka bisa dilakukan
	// dengan fungsi Unscoped()
}

type Role string

// kita perlu menjelaskan Go jika bagian "Role" harus berbentuk string
// selanjutnya kita perlu menentukan role apa saja yg ada di project kita ini

const (
	// kita menggunakan "const" untuk menandakan klo hanya ada "user" dan "admin" saja
	Enduser Role = "user"
	Admin   Role = "admin"
)

func (p *Role) Scan(value any) error {
	// kita membuat function yg dipanggil otomatis saat gorm/ database
	// membaca data dari database dan mau memasukkan ke struct Go
	// func -> mendeklarasikan bahwa ini adalah sebuah function
	// (p *Role) -> adalah variabel penerima yg menmpel pada tipe "Role"
	// dan "p" adalah nama variabel penerima nya,
	// sedangkan "*Role" adalah mengirim pointer ke "Role"
	// tujuannya mengubah nilai asli dari "Role"
	// karena jika hanya menggunakan "Role" tanpa pointer,
	// perubahan hanya terjadi di dalam fungsi dan di luar fungsi data akan tetap menjadi mentah
	// sehingga ketika "Role" tanpa pointer di luar fungsi belum tentu bisa terbaca
	// "Scan" adalah nama fungsi nya. Nama wajib sama karena akan mencocokkan dengan interface sql.Scanner
	// kemudian membuat parameter penerima "value", yg menerima data dalam bentuk apa saja "any"
	// fungsi "Scan" mengembalikan error jika proses scan gagal
	*p = Role(value.([]byte))
	// *p isi/ nilai ini ditujukan kepada pointer "p"
	// "Role(value.([]byte))" adalah konversi tipe
	// hasil "value.[]byte" akan dikonversi ke tipe "Role"
	// kita perlu konversi karena tipe "Role" harus bernilai string
	// "value.[]byte" kita menganggap "value" memiliki nilai yg tidak bisa dipastikan "any"
	// sehingga kita perlu memaksa Go memperlakukan sebagai "[]byte"
	// "[]byte" atau slice of byte adalah representasi teks mentah yg biasa dikirim ke database
	return nil
	// hasil akhir nya adalah mengembalikan tidak ada error
	// dan proses scan berhasil
	// database menerima data mentah ([]byte) -> fungsi "Scan" menangkap melalui "value"
	// -> data di-konvert menjadi bentuk tipe "Role"
	// -> dimasukkan ke alamat asli variabel "p" melalui pointer -> proses scan berhasil, tanpa error
}

func (p Role) Value() (driver.Value, error) {
	// func -> deklarasi fungsi
	// (p Role) -> adalah variabel penerima, kali ini tidak memerlukan "*" pointer
	// artinya menerima "Role" biasa yg memiliki nilai asli yg tidak ada konvert
	// dan hanya membaca saja
	// berbeda dengan "Scan" yg bertugas untuk mengisi dan mengubah nilai asli
	// Value() -> ini adalah nama fungsinya dengan "()" tanpa isi
	// karena tidak ada parameter yg diterima
	// (driver.Value, error) -> fungsi "Value" ini mengembalikan dua nilai,
	// yaitu "driver.Value" -> "driver" nilai yg siap dikirim ke database
	// dan "error" jika ada masalah saat konversi
	return string(p), nil
	// mengembalikan nilai variabel "p" dalam bentuk string karena "Role" berbentuk string
	// dan tidak ada error "nil"
}

// Keberadaan "Scan" sebagai pintu masuk dari database ke aplikasi Go
// Keberadaan "Value" sebagai pintu keluar dari aplikasi Go ke database
// Jika salah satu ada yg tidak ada, misalnya hanya ada fungsi "Value"
// yg hanya bisa menulis saja
// ketika ingin membaca kembali (SELECT) data itu dari database,
// Go akan kebingungan cara konvert "[]byte" ke tipe "Role" yg harus berbentuk string
// sehingga terjadi error
