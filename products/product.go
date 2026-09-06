package products

import (
	"finalissima_e_commerce_rest_api/categories"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	NameProduct string              `json:"name_product"`
	CategoryID  uint                `json:"category_id"`
	Category    categories.Category `json:"category"`
	Description string              `json:"description"`
	Company     string              `json:"company"`
	Barcode     string              `json:"barcode"`
	Weight      uint                `json:"weight"`
	ExpiredDate string              `json:"expired_date"`
	Price       int64               `json:"price"`
	Stock       uint                `json:"stock"`
	ImageLink   string              `json:"image_link"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}

type ProductRequest struct {
	NameProduct string `form:"name_product" validate:"required"`
	CategoryID  uint   `form:"category_id" validate:"required"`
	Description string `form:"description" validate:"required"`
	Company     string `form:"company" validate:"required"`
	Barcode     string `form:"barcode"`
	Weight      uint   `form:"weight" validate:"required"`
	ExpiredDate string `form:"expired_date"`
	Price       int64  `form:"price" validate:"required"`
	Stock       uint   `form:"stock" validate:"required"`
	ImageLink   string
	File        *multipart.FileHeader
}
