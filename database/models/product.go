package models

import (
	"time"

	"gorm.io/gorm"
)

// setelah kita menyelesaikan isi folder "category",
// dalam e-commerce memerlukan adanya fitur produk maka kita buat models "Product" terlebih dahulu
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	NameProduct string         `json:"name_product" gorm:"unique"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `json:"category"`
	Description string         `json:"description"`
	Company     string         `json:"company"`
	Barcode     string         `json:"barcode" gorm:"unique"`
	Weight      uint           `json:"weight"`
	ExpiredDate string         `json:"expired_date"`
	Price       int64          `json:"price"`
	Stock       uint           `json:"stock"`
	ImageLink   string         `json:"image_link"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
