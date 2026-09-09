package api

import (
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/auth"
	"finalissima_e_commerce_rest_api/categories"
	"finalissima_e_commerce_rest_api/package/rajaongkir"
	"finalissima_e_commerce_rest_api/products"
	"finalissima_e_commerce_rest_api/users"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary, jwtConfig middlewares.JWTConfig) *echo.Echo {
	var (
		e               = echo.New()
		authService     = auth.New(repository)
		userService     = users.New(repository)
		categoryService = categories.New(repository)
		productService  = products.New(repository)
		roService       = rajaongkir.InitService()
	)
}
