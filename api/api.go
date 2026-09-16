package api

import (
	"finalissima_e_commerce_rest_api/api/handlers"
	"finalissima_e_commerce_rest_api/api/middlewares"
	"finalissima_e_commerce_rest_api/auth"
	"finalissima_e_commerce_rest_api/categories"
	"finalissima_e_commerce_rest_api/package/ai"
	"finalissima_e_commerce_rest_api/package/constant"
	"finalissima_e_commerce_rest_api/package/fileupload"
	"finalissima_e_commerce_rest_api/products"
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
		categoryService   = categories.New(repository)
		productService    = products.New(repository)
		aiSeervice        = ai.InitService()
		uploader          = &fileupload.CloudinaryUploader{Cld: cld}
		authHandler       = handlers.NewAuth(authService, jwtConfig)
		categoriesHandler = handlers.NewCategories(categoryService)
		productsHandler   = handlers.NewProducts(productService, uploader, aiSeervice)
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

	categoryRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	categoryRoutes.POST("/categories", categoriesHandler.CreateCategory, middlewares.VerifyAdmin, middlewares.ValidateBody(&categories.CategoryRequest{}))
	categoryRoutes.GET("/categories/:id", categoriesHandler.GetByID)
	categoryRoutes.GET("/categories", categoriesHandler.GetAllCategories)
	categoryRoutes.PUT("/categories/:id", categoriesHandler.UpdateCategory, middlewares.VerifyAdmin, middlewares.ValidateBody(&categories.CategoryRequest{}))
	categoryRoutes.DELETE("/categories/:id", categoriesHandler.DeleteCategoryByID, middlewares.VerifyAdmin)

	productRoutes := e.Group(constant.API_V1_PREFIX, echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	productRoutes.POST("/product", productsHandler.CreateProduct, middlewares.VerifyAdmin, middlewares.ValidateBody(&products.Product{}))
	productRoutes.GET("/product/:id", productsHandler.GetProductByID)
	productRoutes.GET("/products/category/:id", productsHandler.GetProductsByCategory)
	productRoutes.GET("/products", productsHandler.GetAllProducts)
	productRoutes.PUT("/product/:id", productsHandler.UpdateProduct, middlewares.VerifyAdmin, middlewares.ValidateBody(&products.ProductRequest{}))

	return e
}
