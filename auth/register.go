package auth

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type register struct {
	repository *gorm.DB
}

func (r register) RegisterUser(ctx context.Context, registerRequest *RegisterRequest) (User, error) {
	password, err := utils.GeneratePassword(registerRequest.Password)
	// ini adalah bagian password di-generate menjadi token bcrypt
	if err != nil {
		return User{}, err
	}

	user := models.User{
		Username:    registerRequest.Username,
		Email:       registerRequest.Email,
		Password:    string(password),
		PhoneNumber: registerRequest.PhoneNumber,
		Address:     registerRequest.Address,
		ProvinceID:  registerRequest.ProvinceID,
		CityID:      registerRequest.CityID,
		DistrictID:  registerRequest.DistrictID,
		Role:        models.Enduser,
	}
	result := r.repository.WithContext(ctx).Create(&user)
	// kita menggunakan pointer "&" memiliki maksud user mengisi datanya
	if err := result.Error; err != nil {
		return User{}, err
	}

	record := new(User)
	if err := result.WithContext(ctx).Last(record).Error; err != nil {
		// ini adalah proses data user baru disimpan ke dalam database
		return User{}, err
	}

	return *record, nil
}
