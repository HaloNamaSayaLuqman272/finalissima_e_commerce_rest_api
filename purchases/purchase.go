package purchases

import (
	"finalissima_e_commerce_rest_api/database/models"
	"finalissima_e_commerce_rest_api/products"
	"os/user"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID         uint                  `json:"id" gorm:"primaryKey"`
	UserID     uint                  `json:"user_id"`
	User       user.User             `json:"user"`
	ProductID  uint                  `json:"product_id"`
	Product    products.Product      `json:"product"`
	Price      float64               `json:"price"`
	Quantity   uint                  `json:"quantity"`
	Amount     float64               `json:"amount"`
	Fee        float64               `json:"fee"`
	Courier    string                `json:"courier"`
	Status     models.PurchaseStatus `json:"status"`
	ReceivedAt time.Time             `json:"received_at"`
	CreatedAt  time.Time             `json:"created_at"`
	DeletedAt  *gorm.DB              `json:"deleted_at" gorm:"index"`
}
