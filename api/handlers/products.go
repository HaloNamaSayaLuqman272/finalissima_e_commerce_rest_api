package handlers

import (
	"finalissima_e_commerce_rest_api/package/ai"
	"finalissima_e_commerce_rest_api/package/fileupload"
	"finalissima_e_commerce_rest_api/products"

	"github.com/labstack/echo/v5"
)

type Products struct {
	products       products.Service
	recommendation ai.Service
	uploader       fileupload.Uploader
}

func (p Products) CreateProduct(ctx *echo.Context) error {
	productRequest := ctx.Get("validatedBody").(*products.Product)
}
