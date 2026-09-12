package purchases

import (
	"finalissima_e_commerce_rest_api/database/models"
	"finalissima_e_commerce_rest_api/products"
	"os/user"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID                uint                  `json:"id" gorm:"primaryKey"`
	UserID            uint                  `json:"user_id"`
	User              user.User             `json:"user"`
	ProductID         uint                  `json:"product_id"`
	Product           products.Product      `json:"product"`
	Price             float64               `json:"price"`
	Quantity          uint                  `json:"quantity"`
	Weight            uint                  `json:"weight"`
	Amount            float64               `json:"amount"`
	DestinationID     uint                  `json:"destination_id"`
	Fee               float64               `json:"fee"`
	Courier           string                `json:"courier"`
	Status            models.PurchaseStatus `json:"status"`
	EstimatedDelivery string                `json:"estmated_delivery"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         *time.Time            `json:"updated_at"`
	ReceivedAt        *time.Time            `json:"received_at"`
	DeletedAt         *gorm.DB              `json:"deleted_at" gorm:"index"`
}

type PurchaseRequest struct {
	Products      []uint `json:"products" validate:"required"`
	Courier       string `json:"courier" validate:"required,validCourier"`
	DestinationID uint   `json:"destination_id" validate:"required"`
	Weight        uint   `json:"weight" validate:"required,min=1"`
	UserID        uint
	Fee           float64
	ReceivedTime  *time.Time
}

type UpdatePurchaseRequest struct {
	Status string `json:"status" validate:"required"`
}
