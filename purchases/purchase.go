package purchases

import (
	"finalissima_e_commerce_rest_api/database/models"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID                  uint                  `json:"id" gorm:"primaryKey"`
	UserID              uint                  `json:"user_id"`
	User                models.User           `json:"user"`
	Items               []models.PurchaseItem `json:"items" gorm:"foreignKey:PurchaseID"`
	TotalWeight         uint                  `json:"total_weight"`
	Amount              float64               `json:"amount"`
	DestinationID       uint                  `json:"destination_id"`
	Fee                 float64               `json:"fee"`
	Courier             string                `json:"courier"`
	Status              models.PurchaseStatus `json:"status"`
	PaidAt              *time.Time            `json:"paid_at"`
	TrackingNumber      string                `json:"tracking_number"`
	ShippedAt           *time.Time            `json:"shipped_at"`
	CreatedAt           time.Time             `json:"created_at"`
	UpdatedAt           *time.Time            `json:"updated_at"`
	EstimatedDelivery   uint                  `json:"estimated_delivery"`
	EstimatedReceivedAt *time.Time            `json:"estimated_received_at"`
	ReceivedAt          *time.Time            `json:"received_at"`
	DeletedAt           gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
}

type PurchaseItemOrder struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required,min=1"`
}

type PurchaseOrder struct {
	Items               []PurchaseItemOrder `json:"items" validate:"required,min=1,dive"`
	Courier             string              `json:"courier" validate:"required,validCourier"`
	DestinationID       uint                `json:"destination_id" validate:"required"`
	UserID              uint                `json:"-"`
	TotalWeight         uint                `json:"-"`
	Amount              float64             `json:"-"`
	Fee                 float64             `json:"_"`
	EstimatedDelivery   uint                `json:"-"`
	EstimatedReceivedAt *time.Time          `json:"-"`
}

type UpdatePurchaseOrder struct {
	Status string `json:"status" validate:"required"`
}
