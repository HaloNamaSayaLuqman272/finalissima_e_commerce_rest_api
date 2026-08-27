package users

import (
	"context"

	"gorm.io/gorm"
)

// setelah kita menentukan data user apa saja yg diperlukan untuk project ini
// kita perlu membuat alur proses user mengisi data yg sudah ditentukan
type get struct {
	// membuat type struct "get" rencnanya akan digunakan untuk membuat function
	repository *gorm.DB
	// kemudian di dalam type struct "get" di dalamnya kita kita isi dengan
	// repository yg terhubung dengan gorm.DB
	// tujuan nya agar database siap dipakai kapan saja
	// dan tidak perlu mengulang mengirim atau mengoper parameter dalam pemanggilan fungsi
}

func (g get) GetProfile(ctx context.Context, userID uint) (User, error) {
	// kita membuat fungsi untuk menjalankan bagaimana alur proses "GetProfile" berlangsung
	// kita membuat "g" sebagai variabel penerima dari type struct "get"
	// () setelah "GetProfile" adalah parameter list sebagai jalur masuk
	// untuk memasukkan data ke dalam fungsi
	// selanjutnya mengisi parameter list tersebut dengan
	// membuat parameter "ctx" dengan tipe data context.Context di dalamnya
	// untuk menjaga lifecycle sebuah request
	// dan memabuat parameter "userID" dengan tipe data uint
	// untuk mengambil data dari ID user mana yg akan diambil
	// (User, error) -> artinya mengembalikan data User atau error
	user := new(User)
	// dalam fungsi GetProfile kita membuat nama variabel baru "user"
	// yg akan menampung isi data "User" yg diinput
	err := g.repository.WithContext(ctx).
		// "err" adalah variabel yg akan digunakan untuk menampung hasil ".Error"
		// "g" adalah objek dari repository
		// "WithContext(ctx)" artinya mengirim parameter context.Context
		// bertujuan untuk mengatur timeout pengisian atau membatalkan proses query jika koneksi pengguna terputus
		First(user, "id = ?", userID).
		// menggunakan query SQL "First" dari gorm untuk mengambil satu data pertama
		//  yg cocok dengan kondisi "id = userID",
		// jika telah ketemu maka akan dimasukkan ke dalam variabel pointer "user"
		Error
		// "Error" digunakan untuk mengembil properti "error" dari objek hasil query
		// jika sukses maka akan bernilai "nil" atau bersih dari error
		// jika koneksi putus atau data tidak ketemu maka nilai nya akan berisi "error"

	if err != nil {
		// jika variabel "err" hasilnya tidak sama dengan "nil"
		return User{}, err
		// maka akan kembalikan struct "User{}" kosongan dan menampilkan "error"
	}

	return *user, nil
	// jika berhasil maka kembalikan nilai hasil inputan "user"
	// dan bernilai "nil" bersih dari "error"
}
