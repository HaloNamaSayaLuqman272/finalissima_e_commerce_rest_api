package models

import (
	"time"

	"gorm.io/gorm"
)

// setelah selesai dengan bagian "users", kita akan membuat model untuk "categories"
// kita membuat type struct "Category"
type Category struct {
	ID           uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	NameCategory string          `json:"name_category" gorm:"unique"`
	Description  string          `json:"description"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
