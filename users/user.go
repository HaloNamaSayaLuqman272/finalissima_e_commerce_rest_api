package users

// kita membutuhkan data apa saja yg perlu user isi untuk project kita ini
// karena rencananya akan ada banyak data yg diperlukan,
// maka kita bungkus dalam type struct untuk menampung nama data-data tersebut
type Users struct {
	ID uint `json:"id" form:"-" gorm:"primaryKey;autoIncrement"`
	// primaryKey -> field ini akan dijadikan sebagai kunci utama tabel SQL ||
	// autoIncrement -> perintah untuk database agar mengisi ID otomatis sesuai dengan angka urutan
	// tanda (-) -> abaikan field ini dan user tidak bisa mengisi ID sendiri
	// json -> menyertakan ID agar sistem mengembalikan respon pada frontend jika registrasi user berhasil
	Name string `json:"name"`
}
