package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"time"

	"gorm.io/gorm"
)

type updatepurchasestatus struct {
	repository *gorm.DB
	getbypurchaseid
}

func (u updatepurchasestatus) UpdatePurchaseStatus(ctx context.Context, purchaseRequest *UpdatePurchaseOrder, id uint) (Purchase, error) {
	status := models.PurchaseStatus(purchaseRequest.Status)
	purchase := models.Purchase{
		Status: status,
	}
	if status == models.Paid {
		purchase.PaidAt = time.Now()
	}
}
