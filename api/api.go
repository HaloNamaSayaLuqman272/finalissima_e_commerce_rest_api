package api

import (
	"finalissima_e_commerce_rest_api/api/handlers"
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/auth"
	"finalissima_e_commerce_rest_api/categories"
	"finalissima_e_commerce_rest_api/package/ai"
	"finalissima_e_commerce_rest_api/package/constant"
	"finalissima_e_commerce_rest_api/package/fileupload"
	"finalissima_e_commerce_rest_api/package/rajaongkir"
	"finalissima_e_commerce_rest_api/products"
	"finalissima_e_commerce_rest_api/purchases"
	"finalissima_e_commerce_rest_api/users"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary, jwtConfig middlewares.JWTConfig) *echo.Echo {
	var (
		e                 = echo.New()
		authService       = auth.New(repository)
		userService       = users.New(repository)
		categoryService   = categories.New(repository)
		productService    = products.New(repository)
		roService         = rajaongkir.InitService()
		aiSeervice        = ai.InitService()
		shippingService   = purchases.NewFakeShippingService()
		purchaseService   = purchases.New(repository, productService, roService, shippingService)
		uploader          = &fileupload.CloudinaryUploader{Cld: cld}
		authHandler       = handlers.NewAuth(authService, jwtConfig)
		userHandler       = handlers.NewUsers(userService, uploader)
		categoriesHandler = handlers.NewCategories(categoryService)
		productsHandler   = handlers.NewProducts(productService, uploader, aiSeervice)
		purchaseHandler   = handlers.NewPurchases(purchaseService, userService, roService)
	)

	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	logger, _ := zap.NewProduction()
	loggerConfig := middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}

	logMiddleware := middlewares.LoggerConfig{Config: loggerConfig}
	jwtMiddleware := jwtConfig.Init()

	e.Use(logMiddleware.Init())

	authRoutes := e.Group(fmt.Sprintf("%s/auth", constant.API_V1_PREFIX))

	authRoutes.POST("/register", authHandler.RegisterUser)
	authRoutes.POST("/login", authHandler.LoginUser)

	userRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	userRoutes.GET("/profile", userHandler.GetProfile)
	userRoutes.PATCH("/profile/edit", userHandler.UpdateProfile)

	categoryRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	categoryRoutes.POST("/categories", categoriesHandler.CreateCategory, middlewares.VerifyAdmin, middlewares.ValidateBody(&categories.CategoryRequest{}))
	categoryRoutes.GET("/categories/:id", categoriesHandler.GetByID)
	categoryRoutes.GET("/categories", categoriesHandler.GetAllCategories)
	categoryRoutes.PUT("/categories/:id", categoriesHandler.UpdateCategory, middlewares.VerifyAdmin, middlewares.ValidateBody(&categories.CategoryRequest{}))
	categoryRoutes.DELETE("/categories/:id", categoriesHandler.DeleteCategoryByID, middlewares.VerifyAdmin)

	productRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	productRoutes.POST("/products", productsHandler.CreateProduct, middlewares.VerifyAdmin, middlewares.ValidateBody(&products.ProductRequest{}))
	productRoutes.GET("/products/:id", productsHandler.GetProductByID)
	productRoutes.GET("/products/category/:id", productsHandler.GetProductsByCategory)
	productRoutes.POST("/products/recommendation", productsHandler.GetProductRecommendation, middlewares.ValidateBody(&ai.ProductRecommendationRequest{}))
	productRoutes.GET("/products", productsHandler.GetAllProducts)
	productRoutes.PUT("/products/:id", productsHandler.UpdateProduct, middlewares.VerifyAdmin, middlewares.ValidateBody(&products.ProductRequest{}))
	productRoutes.DELETE("/products/:id", productsHandler.DeleteProduct, middlewares.VerifyAdmin)

	webhookRoutes := e.Group(fmt.Sprintf("%s/webhooks", constant.API_V1_PREFIX))
	webhookRoutes.POST("/purchases/:id/payment-success", purchaseHandler.ConfirmPayment)

	purchaseRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	purchaseRoutes.POST("/purchases", purchaseHandler.CreatePurchase, middlewares.ValidateBody(&purchases.PurchaseOrder{}))
	purchaseRoutes.GET("/purchases/my", purchaseHandler.GetMyPurchases)
	purchaseRoutes.GET("/purchases/user/:id", purchaseHandler.GetPurchasesByUserID, middlewares.VerifyAdmin)
	purchaseRoutes.GET("/purchases/:id", purchaseHandler.GetByPurchaseID, middlewares.VerifyAdmin)
	purchaseRoutes.GET("/purchases", purchaseHandler.GetAllPurchases, middlewares.VerifyAdmin)
	purchaseRoutes.PATCH("/purchases/:id", purchaseHandler.UpdatePurchaseStatus, middlewares.VerifyAdmin, middlewares.ValidateBody(&purchases.UpdatePurchaseOrder{}))
	purchaseRoutes.PATCH("/purchases/:id/cancel", purchaseHandler.CancelPurchase)
	purchaseRoutes.DELETE("/purchases/:id", purchaseHandler.DeletePurchase, middlewares.VerifyAdmin)

	return e
}
