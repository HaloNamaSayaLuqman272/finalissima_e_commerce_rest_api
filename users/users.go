package users

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

// disini kita memalakukan copy paste dari data "user.go" dari package models
type User struct {
	ID             uint           `json:"id" form:"-" gorm:"primaryKey;autoIncrement"`
	Username       string         `json:"username" gorm:"unique"`
	Email          string         `json:"email" gorm:"unique"`
	Password       string         `json:"password"`
	PhoneNumber    string         `json:"phone_number" gorm:"uniqueIndex"`
	Address        string         `json:"address" gorm:"type:text"`
	ProvinceID     uint           `json:"province_id"`
	CityID         uint           `json:"city_id"`
	DistrictID     uint           `json:"district_id"`
	ProfilePicture string         `json:"profile_picture"`
	Role           string         `json:"role"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// sekarang anggaplah registrasi sudah dibuat dan berhasil disimpan,
// suatu saat user mencoba memperbarui data,
// maka kita perlu menyediakan type struct untuk UpdateProfileRequest
type EditProfileRequest struct {
	Username string `form:"username" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"min=8,containsSpecialCharacter,containsNumber"`
	// berbeda dengan username dan email yg menggunakan validate required,
	// password membutuhkan validasi yg mengandung angka dan karakter spesial
	// tujuan nya agar non-user tidak bisa dengan mudah menebak password baru
	PhoneNumber    string `form:"phone_number" validate:"required,min=10,containsNumberOnly"`
	Address        string `form:"address" validate:"required,containsNumber"`
	ProvinceID     uint   `form:"province_id" validate:"required"`
	DistrictID     uint   `form:"district_id" validate:"required"`
	ProfilePicture string
	File           *multipart.FileHeader
	// untuk file kita menggunakan multipart.FileHeader
	// dengan tujuan menampung informasi dan metadata dari request HTTP yg mengandung
	// nama file dan ukuran file, serta tipe content
}
