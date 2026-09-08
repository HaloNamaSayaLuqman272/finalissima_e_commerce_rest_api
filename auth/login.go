package auth

import (
	"context"
	"finalissima_e_commerce_rest_api/package/utils"

	"gorm.io/gorm"
)

type login struct {
	repository *gorm.DB
}

func (l login) LoginUser(ctx context.Context, loginRequest *LoginRequest) (User, error) {
	user := new(User)
	if err := l.repository.WithContext(ctx).First(user, "email = ?", loginRequest.Email).Error; err != nil {
		return User{}, err
	}

	err := utils.ComparePassword(user.Password, loginRequest.Password)
	if err != nil {
		return User{}, err
	}

	return *user, nil
}
