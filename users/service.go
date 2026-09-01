package users

import "context"

// pada bagian "service.go" kita membuat type yg berbentuk interface,
// rencananya interface ini akan menjadi jalan penghubung "User"
// dengan "handler" dan "repository"
// dan kita akan menggunakan "mockery" untuk membuat testing
// kita akan menamakan type interface ini dengan nama "Service"
type Service interface {
	// dalam type interface "Service" kita mendeklarasikan 2 buah method yg sudah kita buat sebelumnya
	GetProfile(ctx context.Context, userID uint) (User, error)
	// method pertama adalah GetProfile, yg bertujuan untuk mengambil data user berdasarkan ID
	// kita membuat parameter "context.Context" dengan dengan nama "ctx"
	// dan parameter "userID" dengan tipe data uint, yg akan menjadi ID user yg ingin kita ambil datanya
	// method ini akan mengembalikan ke struct "User" dan error
	UpdateProfile(ctx context.Context, editReq *EditProfileRequest, id uint) (User, error)
	// method kedua adalah UpdateProfile, yg bertujuan untuk mengubah data user berdasarkan ID
	// kita membuat parameter "context.Context" dengan dengan nama "ctx"
	// dan parameter "editReq" dengan tipe data pointer ke struct "EditProfileRequest"
	// yg akan menjadi data baru yg ingin kita ubah
	// dan parameter "id" dengan tipe data uint, yg akan menjadi ID user yg ingin kita ubah datanya
	// method ini akan mengembalikan ke struct "User" dan error
	// kita mengembalikan ke struct "User" tanpa pointer karena
	// kita ingin mengembalikan data baru yg sudah diubah
}

// selanjutnya kita membuat type struct "service" yg akan mengimplementasikan interface "Service"
type service struct {
	// membuat type struct "service" rencnanya akan digunakan untuk membuat function
	get
	// kita memanggil type struct "get" yg sudah kita buat sebelumnya
	update
	// kita memanggil type struct "update" yg sudah kita buat sebelumnya
	// ke dalam type struct "service" ini, agar kita bisa mengakses method yg ada di dalamnya
}

// selanjutnya kita membutuhkan sesuatu untuk meengecek apakah "service"
// sudah mengimplementasikan interface "Service" atau belum
// sehingga kita perlu membuat variabel underscore "_"
// yg akan menampung hasil implementasi interface "Service" oleh type struct "service"
var _ Service = (*service)(nil)

// kita alih-alih memberi nama variabel, kita malah menggunakan "_"
// karena kita untuk selanjutnya tidak akan menggunakan variabel ini,
// kita hanya ingin mengecek apakah "service" sudah mengimplementasikan interface "Service" atau belum
// seandainya kita memberi nama variabel namun kita tidak menggunakannya pada proses selanjutnya,
// maka Go akan memunculkan error "declared and not used"
// selanjutnya variabel "_" memakai tipe data interface "Service"
// "(*service)" kita menggunakan pointer "*"
// agar bisa memastikan pengecekan terhadap method "service" sudah dilakukan dengan benar
// "nil" kita menggunakan "nil" karena kita tidak ingin menginisialisasi variabel ini dengan nilai apapun
// seandainya kita tidak menggunakan "nil" maka Go akan memunculkan error "missing argument for nil"
