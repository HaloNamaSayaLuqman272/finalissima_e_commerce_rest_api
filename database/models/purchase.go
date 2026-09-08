package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID         uint           `gorm:"primaryKey"`
	UserID     uint           `json:"user_id"`
	User       User           `json:"user"`
	ProductID  uint           `json:"product_id"`
	Product    Product        `json:"product"`
	Price      float64        `json:"price"`
	Quantity   uint           `json:"quantity"`
	Amount     float64        `json:"amount"`
	Fee        float64        `json:"fee"`
	Courier    string         `json:"courier"`
	Status     PurchaseStatus `json:"status"`
	ReceivedAt time.Time      `json:"received_at"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  *gorm.DB       `json:"deleted_at"`
}

type PurchaseStatus string

const (
	Pending    PurchaseStatus = "pending"
	Paid       PurchaseStatus = "paid"
	OnDelivery PurchaseStatus = "on_delivery"
	Cancelled  PurchaseStatus = "cancelled"
	Received   PurchaseStatus = "received"
)

func (p *PurchaseStatus) Scan(value any) error {
	// kita mendeklarasikan sebuah fungsi dan bernama "Scan",
	// memiliki variabel penerima "p" dengan tipe data pointer ke "PurchaseStatus"
	// fungsi "Scan" memiliki parameter penerima "value" tipe data "any"
	// dan mengembalikan "error" jika terjadi masalah pada proses "Scan"
	*p = PurchaseStatus(value.([]byte))
	return nil
}

func (p PurchaseStatus) Value() (driver.Value, error) {
	return string(p), nil
}

// penjelasan fungsi "Scan()" dan "Value()" sama dengan penjelasan pada
// package "models" file "user.go"
