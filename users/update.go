package users

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"

	"gorm.io/gorm"
)

// setelah sebelumnya kita telah menentukan data apa saja yg diperlukan pada "User"
// dan kita telah membuat alur bagaimana "user" mengisi data apa saja yg sudah kita tentukan
// maka disini kita membuat fungsi untuk alur user jika ingin melakukan update data
// kita buat type struct "update" terlebih dahulu
type update struct {
	repository *gorm.DB
	// kita membutuhkan repository dari "gorm" sama seperti pada type struct "get" sebelumnya
	get
	// pada type struct "update" kita mengambil field type struct "get"
	// artinya type struct "update" bisa mendapatkan akses seluruh isi,
	// method, data, dan sebagainya yg dimiliki oleh type struct "get"
}

func (u update) UpdateProfile(ctx context.Context, editProfileRequest *EditProfileRequest, id uint) (User, error) {
	// kita membuat fungsi bernama "UpdateProfile" untuk menampung data dari
	// fitur "Update" data profile yg kita sediakan
	// seperti pada function "GetProfile",
	// dan membuat variabel penerima bernama "u" bertipe "update"
	// function "UpdateProfile" membutuhkan context.Context untuk menjaga lifecycle dan menentukan waktu maksimal eksekusi request,
	// jika user mengalami disconnect jaringan atau cancelled maka operasi akan dihentikan,
	// dan menyimpan data konseptual selama siklus permintaan berlangsung
	// kita bungkus dan namakan "ctx"
	// selanjutnya kita perlu memanggil field yg akan dieksekusi "EditProfileRequest"
	// dan membuat parameter "editProfileRequest" untuk meampung inputan data yg dimasukkan
	user := models.User{
		// kita membuat variabel baru bernama "user"
		// dengan tipe data "models.User",
		// artinya struct data "User" dan berada pada package "models"
		// selanjutnya kita menyebutkan isi dari struct literal "EditProfileRequest"
		// yg diedit akan ada data apa saja
		Username:       editProfileRequest.Username,
		Email:          editProfileRequest.Email,
		Password:       editProfileRequest.Password,
		PhoneNumber:    editProfileRequest.PhoneNumber,
		Address:        editProfileRequest.Address,
		ProvinceID:     editProfileRequest.ProvinceID,
		DistrictID:     editProfileRequest.DistrictID,
		ProfilePicture: editProfileRequest.ProfilePicture,
	}

	// selanjutnya kita membutuhkan deklarasi variabel baru
	// yg menampung pembaruan "UpdateProfile"
	result := u.repository.WithContext(ctx).
		// kita namakan variabel ini dengan nama "result"
		// variabel "result" ini menampung hasil balikan dari operasi GORM
		// objek yg membawa informasi hasil query, bisa berbentuk error
		// proses berjalannya dengan mengakses field "repository" dari variabel penerima "u"
		// ".WithContext(ctx)" mengikat "ctx" ke query ini,
		// jika request di-cancel atau timeout maka proses update ke database dihentikan
		Where("id = ?", id).
		// "Where" ini adalah bagian filter,
		// karena jika kita tidak menambahakn "Where" maka GORM bisa meng-update semua data id yg ada
		// "("id = ?", id)" adalah klausa SQL dengan placeholder "?"
		// dan "id" adalah nilai yg menggantikan "?"
		// artinya cari baris di database yg kolom "id" nya sama dengan "id" yg diberikan
		Updates(&user)
	// "Updates" adalah method dari GORM untuk mengubah field
	// yg nilai nya tidak kosong pada baris yg sama dengan "Where"
	// mengirim alamat dari variabel "user"
	// kita menggunakan tanda pointer "&" agar GORM membaca membaca struct nya
	// dan mengetahui field mana yg update ini ditujukan

	// selanjutnya kita perlu membuat kondisi pengecekan apakah terjadi error atau tidak
	if err := result.Error; err != nil {
		// "err" kita membuat variabel baru yg gunanya untuk menyimpan hasil operasi "Updates"
		// "result.Error" adalah bagian inisialisasi
		// kita memanggil field "Error" dari variabel "result"
		// variabel "result" bertipe *gorm.DB dan pada Gorm.DB ada field "Error"
		// "err != nil" -> adalah bagian kondisi pengecekan apakah variabel "err"
		// tidak berada kondisi "nil"
		// jika memang variabel "err" tidak bernilai nil, berarti ada error
		// dan akan masuk ke dalam blok "if"
		// kita membutuhkan blok "if"
		// agar variabel "err" ini hanya berjalan di blok "if" saja dan tidak keluar kemana-mana
		return User{}, err
		// mengembalikan struct "User{}" kosongan
		// dan mengembalikan "err" error yg ditangkap agar posisi salahnya diketahui
	}

	record, err := u.GetProfile(ctx, id)
	// selanjutnya kita membuat variabel baru "record" yg menyimpan hasil
	// variabel penerima "u" dari fungsi "UpdateProfile"
	// "ctx" dan "id" adalah dua argumen
	// "ctx" untuk menjaga lifecycle
	// "id" argumen untuk menunjukkan kondisi update di "id" keberapa
	// artinya panggil fungsi "UpdateProfile" untuk mengambil data user yg baru
	// berdasarkan "id" yg sama, simpan hasilnya pada variabel "record"
	// dan jika ada error maka simpan error tersebut ke dalam variabel "err"
	if err != nil {
		// jika variabel "err" adalah tidak bernilai "nil"
		// atau tidak bersih dari error
		return User{}, err
		// maka kembalikan struct "User{}" kosongan dan simpan error tersebut
		// ke dalam variabel "err" yg sudah dideklarasikan di "record, err := ..."
	}

	return record, nil
	// jika proses "Update" berhasil, maka kembalikan hasil data "record"
	// dan "nil" atau bersih dari error
}
