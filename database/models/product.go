package models

import (
	"time"

	"gorm.io/gorm"
)

// setelah kita menyelesaikan isi folder "category",
// dalam e-commerce memerlukan adanya fitur produk maka kita buat models "Product" terlebih dahulu
type Product struct {
	ID          uint            `json:"id" form:"-" gorm:"primaryKey;autoIncrement"`
	NameProduct string          `json:"name_product" gorm:"unique"`
	Description string          `json:"description"`
	Barcode     string          `json:"barcode" gorm:"unique"`
	Company     string          `json:"company"`
	ExpiredDate string          `json:"expired_date"`
	Price       float64         `json:"price"`
	Stock       int             `json:"stock"`
	CategoryID  uint            `json:"category_id"`
	Category    Category        `json:"category"`
	ImageLink   string          `json:"image_link"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
