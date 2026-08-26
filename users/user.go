package users

import (
	"time"

	"gorm.io/gorm"
)

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
type Users struct {
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

// sekarang anggaplah registrasi sudah dibuat dan berhasil disimpan,
// suatu saat user mencoba memperbarui data,
// maka kita perlu menyediakan type struct untuk UpdateProfileRequest
type EditProfileRequest struct {
	Username string `form:"username" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}
