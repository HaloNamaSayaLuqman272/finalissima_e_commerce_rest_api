package handlers

import (
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/auth"
	"finalissima_e_commerce_rest_api/package/dtos"
	"finalissima_e_commerce_rest_api/package/utils"
	"net/http"

	"github.com/labstack/echo/v5"
)

type Auth struct {
	auth      auth.Service
	jwtConfig middlewares.JWTConfig
}

func (a Auth) RegisterUser(ctx *echo.Context) error {
	registerRequest := new(auth.RegisterRequest)
	if err := ctx.Bind(registerRequest); err != nil {
		// ".Bind" adalah method framework "echo" yg bertugas mengubah data mentah dari
		// request menjadi struct Go
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	if err := ctx.Validate(registerRequest); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "validation failed",
			Data:    utils.GetValidationErrorMessage(err.Error()),
		})
	}

	user, err := a.auth.RegisterUser(ctx.Request().Context(), registerRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "user registration failed",
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.Response[any]{
		Status:  "success",
		Message: "user registrated",
		Data:    user,
	})
}

func (a Auth) LoginUser(ctx *echo.Context) error {
	loginRequest := new(auth.LoginRequest)
}
