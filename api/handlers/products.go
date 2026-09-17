package handlers

import (
	"encoding/json"
	"finalissima_e_commerce_rest_api/package/ai"
	"finalissima_e_commerce_rest_api/package/dtos"
	"finalissima_e_commerce_rest_api/package/fileupload"
	"finalissima_e_commerce_rest_api/package/utils"
	"finalissima_e_commerce_rest_api/products"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
)

type Products struct {
	products       products.Service
	recommendation ai.Service
	uploader       fileupload.Uploader
}

func (p Products) CreateProduct(ctx *echo.Context) error {
	productReq := ctx.Get("validatedBody").(*products.ProductRequest)
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "file not found",
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	isFileValid := utils.ValidateFile(ext)
	if !isFileValid {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid file format",
		})
	}

	fileLink, err := p.uploader.UploadFile(ctx.Request().Context(), file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "upload failed",
		})
	}

	productReq.ImageLink = fileLink
	product, err := p.products.CreateProduct(ctx.Request().Context(), productReq)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "upload failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[products.Product]{
		Status:  "success",
		Message: "product created",
		Data:    product,
	})
}

func (p Products) GetProductByID(ctx *echo.Context) error {
	param := ctx.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	product, err := p.products.GetProductByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "product not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[products.Product]{
		Status:  "success",
		Message: "product found",
		Data:    product,
	})
}

func (p Products) GetProductsByCategory(ctx *echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	sort := ctx.QueryParam("sort")
	search := ctx.QueryParam("search")

	categoryIdParam := ctx.Param("id")
	categoryId, err := strconv.ParseUint(categoryIdParam, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid category id",
		})
	}

	pagination := utils.Pagination{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		Search:  search,
		Keyword: "name_product",
	}
	productsData, err := p.products.GetProductByCategory(ctx.Request().Context(), pagination, uint(categoryId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch products by category failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[utils.Pagination]{
		Status:  "success",
		Message: "products by category",
		Data:    productsData,
	})
}

func (p Products) GetProductRecommendation(ctx *echo.Context) error {
	recommendationReq := ctx.Get("validatedBody").(*ai.ProductRecommendationRequest)
	req := ai.ProductRecommendationRequest{
		Quantity: recommendationReq.Quantity,
		Topic:    recommendationReq.Topic,
	}
	res, err := p.recommendation.GetProductRecommendation(req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "get recommendation failed",
		})
	}

	resBody := res.Choices[0].Message.Content
	var recommendations []ai.ProductRecommendationResponse
	err = json.Unmarshal([]byte(resBody), (&recommendations))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "get recommendation failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[[]ai.ProductRecommendationResponse]{
		Status:  "success",
		Message: "product recommendations",
		Data:    recommendations,
	})
}

func (p Products) GetAllProducts(ctx *echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	sort := ctx.QueryParam("sort")
	search := ctx.QueryParam("search")

	pagination := utils.Pagination{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		Search:  search,
		Keyword: "name_product",
	}

	productsData, err := p.products.GetAllProducts(ctx.Request().Context(), pagination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch products failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[utils.Pagination]{
		Status:  "success",
		Message: "all products",
		Data:    productsData,
	})
}

func (p Products) UpdateProduct(ctx *echo.Context) error {
	param := ctx.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	productReq := ctx.Get("validatedBody").(*products.ProductRequest)
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "file not found",
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	isFileValid := utils.ValidateFile(ext)
	if !isFileValid {
		return ctx.JSON(http.StatusUnprocessableEntity, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid file format",
		})
	}

	fileLink, err := p.uploader.UploadFile(ctx.Request().Context(), file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "upload failed",
		})
	}

	productReq.ImageLink = fileLink
	product, err := p.products.UpdateProduct(ctx.Request().Context(), productReq, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "update product failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[products.Product]{
		Status:  "success",
		Message: "product updated",
		Data:    product,
	})
}

func (p Products) DeleteProduct(ctx *echo.Context) error {
	param := ctx.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	err = p.products.DeleteProductByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "delete product failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[products.Product]{
		Status:  "success",
		Message: "product deleted",
	})
}

func NewProducts(products products.Service, uploader fileupload.Uploader, recommendation ai.Service) Products {
	return Products{
		products:       products,
		uploader:       uploader,
		recommendation: recommendation,
	}
}
