package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type Purchase struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id"`
	User          User           `json:"user"`
	ProductID     uint           `json:"product_id"`
	Product       Product        `json:"product"`
	Price         float64        `json:"price"`
	Quantity      uint           `json:"quantity"`
	Weight        uint           `json:"weight"`
	Amount        float64        `json:"amount"`
	DestinationID uint           `json:"destination_id"`
	Fee           float64        `json:"fee"`
	Courier       string         `json:"courier"`
	Status        PurchaseStatus `json:"status"`
	ReceivedAt    *time.Time     `json:"received_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     *time.Time     `json:"updated_at"`
	DeletedAt     *gorm.DB       `json:"deleted_at" gorm:"index"`
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
	// "(p *PurchaseStatus)" kita membuat begini agar "Scan(...)" mengubah nilai dari
	// variabel asli yg memanggilnya "p" dan kita menambah "*" agar fungsi "Scan(...)"
	// mendapatkan alamat memori asli variabel "PurchaseStatus". Saat "Scan(...)" berhasil
	// membaca dari database dan mengisinya ke "p", nilai dari variabel asli "PurchaseStatus"
	// akan ikut berubah dan tersimpan
	// fungsi "Scan" memiliki parameter penerima "value" tipe data "any"
	// dan mengembalikan "error" jika terjadi masalah pada proses "Scan"
	*p = PurchaseStatus(value.([]byte))
	// "value.([]byte)" proses type assertion untuk memastikan bahwa data yg dikirim
	// oleh database bertipe "[]byte"
	// "PurchaseStatus(...)" type conversion mengubah data mentah menjadi tipe "PurchaseStatus"
	// "*" digunakan untuk mengakses nilai asli yg ditujukan pointer tersebut "p"
	// arti code ini perintah simpan hasil konversi tersebut ke dalam variabel asli
	// "PurchaseStatus" yg memanggil fungsi "Scan()"
	return nil
}

func (p PurchaseStatus) Value() (driver.Value, error) {
	// "(p PurchaseStatus)" kita membuat begini karena kita ingin Go membuat salinan
	// "PurchaseStatus" saat fungsi "Scan(...)" dipanggil dan segala perubahan yg dilakukan
	// hanya terjadi di salinan nya saja di fungsi "Value"
	return string(p), nil
	// jika berhasil maka menampilkan perubahan "p" dalam bentuk tipe data "string"
}

// PERBEDAAN FUNGSI "Scan(value any)" DAN "Value()"
// Fungsi "Scan (value any)" digunakan saat kita ingin membaca data dari database dan
// memasukkannya kembali ke dalam tipe data kustom Go
// Fungsi "Value()" digunakan untuk konversi data kustom Go menjadi tipe data yg bisa
// disimpan dan dipahami oleh driver database
