package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"fmt"

	"gorm.io/gorm"
)

type cancelpurchase struct {
	repository      *gorm.DB
	getbypurchaseid getbypurchaseid
	shippingService ShippingService
}

func NewCancelPurchase(repository *gorm.DB, getbypurchaseid getbypurchaseid, shippingService ShippingService) cancelpurchase {
	return cancelpurchase{
		repository:      repository,
		getbypurchaseid: getbypurchaseid,
		shippingService: shippingService,
	}
}

func (c cancelpurchase) CancelPurchase(ctx context.Context, id uint, requestingUserID uint, isAdmin bool) (Purchase, error) {
	var purchase models.Purchase
	if err := c.repository.WithContext(ctx).First(&purchase, id).Error; err != nil {
		return Purchase{}, err
	}

	if !isAdmin && purchase.UserID != requestingUserID {
		return Purchase{}, fmt.Errorf("you are not allowed to cancel this purchase")
	}

	if purchase.Status == models.Received {
		return Purchase{}, fmt.Errorf("purchase that have already been received cannot be cancelled")
	}

	if purchase.Status == models.Cancelled {
		return Purchase{}, fmt.Errorf("the purchase was previously cancelled")
	}

	if purchase.Status == models.OnDelivery && purchase.TrackingNumber != "" {
		if err := c.shippingService.CancelShipmentOrder(ctx, purchase.TrackingNumber); err != nil {
			return Purchase{}, fmt.Errorf("failed to cancel the shipping resi: %w", err)
		}
	}

	result := c.repository.WithContext(ctx).Model(&models.Purchase{}).Where("id = ?", id).Update("status", models.Cancelled)
	if result.Error != nil {
		return Purchase{}, result.Error
	}

	purchaseRecord, err := c.getbypurchaseid.GetByPurchaseID(ctx, id)
	if err != nil {
		return Purchase{}, err
	}

	return purchaseRecord, nil
}
