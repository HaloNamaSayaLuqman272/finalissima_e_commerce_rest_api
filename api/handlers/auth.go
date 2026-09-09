package handlers

import (
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/auth"
)

type Auth struct {
	auth      auth.Service
	jwtConfig middlewares.JWTConfig
}
