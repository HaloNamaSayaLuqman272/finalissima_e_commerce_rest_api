package handlers

import (
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/database/models"
	"finalissima_e_commerce_rest_api/package/dtos"
	"finalissima_e_commerce_rest_api/package/rajaongkir"
	"finalissima_e_commerce_rest_api/package/utils"
	"finalissima_e_commerce_rest_api/purchases"
	"finalissima_e_commerce_rest_api/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Purchases struct {
	purchases purchases.Service
	users     users.Service
	roService rajaongkir.Service
}

func (p Purchases) CreatePurchase(ctx *echo.Context) error {
	purchaseReq := ctx.Get("validatedBody").(*purchases.PurchaseOrder)

	userData := ctx.Get("userData").(*middlewares.JWTCustomsClaims)
	user, err := p.users.GetProfile(ctx.Request().Context(), uint(userData.ID))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "user not found",
		})
	}

	purchaseReq.UserID = uint(user.ID)

	createdPurchase, err := p.purchases.CreatePurchase(ctx.Request().Context(), purchaseReq)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.Response[purchases.Purchase]{
		Status:  "success",
		Message: "purchase created successfully",
		Data:    createdPurchase,
	})
}

func (p Purchases) GetByPurchaseID(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid id",
		})
	}

	purchase, err := p.purchases.GetByPurchaseID(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dtos.Response[any]{
			Status:  "failed",
			Message: "purchase not found",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[any]{
		Status:  "failed",
		Message: "purchase found",
		Data:    purchase,
	})
}

func (p Purchases) GetMyPurchases(ctx *echo.Context) error {
	userData := ctx.Get("userData").(*middlewares.JWTCustomsClaims)

	purchasseData, err := p.purchases.GetPurchaseByUserID(ctx.Request().Context(), uint(userData.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch purchases failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[[]purchases.Purchase]{
		Status:  "success",
		Message: "my purchases",
		Data:    purchasseData,
	})
}

func (p Purchases) GetPurchasesByUserID(ctx *echo.Context) error {
	param := ctx.Param("id")

	userID, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid user id",
		})
	}

	purchasesData, err := p.purchases.GetPurchaseByUserID(ctx.Request().Context(), uint(userID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch purchases failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[[]purchases.Purchase]{
		Status:  "success",
		Message: "purchases by user id",
		Data:    purchasesData,
	})
}

func (p Purchases) GetAllPurchases(ctx *echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	sort := ctx.QueryParam("sort")
	search := ctx.QueryParam("search")

	pagination := utils.Pagination{
		Page:    page,
		Limit:   limit,
		Sort:    sort,
		Search:  search,
		Keyword: "courier",
	}

	purchasesData, err := p.purchases.GetAllPurchases(ctx.Request().Context(), pagination)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.Response[any]{
			Status:  "failed",
			Message: "fetch purchases failed",
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[[]purchases.Purchase]{
		Status:  "success",
		Message: "all purchases",
		Data:    purchasesData,
	})
}

func (p Purchases) ConfirmPayment(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid purchase id",
		})
	}

	updatedPurchase, err := p.purchases.ConfirmPayment(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[purchases.Purchase]{
		Status:  "success",
		Message: "purchase payment confirmed",
		Data:    updatedPurchase,
	})
}

func (p Purchases) UpdatePurchaseStatus(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid purchase id",
		})
	}

	updateReq := ctx.Get("validatedBody").(*purchases.UpdatePurchaseOrder)
	updatedPurchase, err := p.purchases.UpdatePurchaseStatus(ctx.Request().Context(), updateReq, uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[purchases.Purchase]{
		Status:  "success",
		Message: "purchase status updated",
		Data:    updatedPurchase,
	})
}

func (p Purchases) CancelPurchase(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid purchase id",
		})
	}

	userData := ctx.Get("userData").(*middlewares.JWTCustomsClaims)
	isAdmin := userData.Role == models.Admin

	cancelledPurchase, err := p.purchases.CancelPurchase(ctx.Request().Context(), uint(id), uint(userData.ID), isAdmin)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[purchases.Purchase]{
		Status:  "success",
		Message: "purchase cancelled",
		Data:    cancelledPurchase,
	})
}

func (p Purchases) DeletePurchase(ctx *echo.Context) error {
	param := ctx.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: "invalid purchase id",
		})
	}

	err = p.purchases.DeletePurchase(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.Response[any]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dtos.Response[any]{
		Status:  "success",
		Message: "purchase deleted successfully",
	})
}

func NewPurchases(purchases purchases.Service, users users.Service, roService rajaongkir.Service) Purchases {
	return Purchases{
		purchases: purchases,
		users:     users,
		roService: roService,
	}
}
