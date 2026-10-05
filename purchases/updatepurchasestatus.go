package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type updatepurchasestatus struct {
	repository      *gorm.DB
	getbypurchaseid getbypurchaseid
	shippingService ShippingService
}

func NewUpdatePurchaseStatus(repository *gorm.DB, getbypurchaseid getbypurchaseid, shippingService ShippingService) updatepurchasestatus {
	return updatepurchasestatus{
		repository:      repository,
		getbypurchaseid: getbypurchaseid,
		shippingService: shippingService,
	}
}

func (u updatepurchasestatus) processShipment(ctx context.Context, id uint) error {
	fmt.Println("STEP 1: processShipment begin, id = ", id)

	var purchase models.Purchase
	if err := u.repository.WithContext(ctx).First(&purchase, id).Error; err != nil {
		return err
	}
	fmt.Println("STEP 2: fetch purchase data success")

	shipment, err := u.shippingService.CreateShipmentOrder(ctx, ShippingRequest{
		OrderID:       purchase.ID,
		DestinationID: purchase.DestinationID,
		Courier:       purchase.Courier,
		Weight:        purchase.TotalWeight,
	})
	if err != nil {
		fmt.Println("STEP 3 ERROR ", err)
		return err
	}
	fmt.Println("STEP 3: CreateShipmentOrder success, tracking =", shipment.TrackingNumber)

	now := time.Now()
	result := u.repository.WithContext(ctx).Model(&models.Purchase{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":          models.OnDelivery,
		"tracking_number": shipment.TrackingNumber,
		"shipped_at":      now,
	})
	fmt.Println("STEP 4: update kedua successfully, error =", result.Error)

	return result.Error
}

func (u updatepurchasestatus) ConfirmPayment(ctx context.Context, id uint) (Purchase, error) {
	result := u.repository.WithContext(ctx).Model(&models.Purchase{}).Where("id = ? AND status = ?", id, models.Pending).Updates(map[string]interface{}{
		"status":  models.Paid,
		"paid_at": time.Now(),
	})

	if result.Error != nil {
		return Purchase{}, result.Error
	}
	if result.RowsAffected == 0 {
		return Purchase{}, fmt.Errorf("purchase already paid or not found")
	}

	if err := u.processShipment(ctx, id); err != nil {
		return Purchase{}, fmt.Errorf("recorded payment, but failed to create shipment: %w", err)
	}

	purchaseRecord, err := u.getbypurchaseid.GetByPurchaseID(ctx, id)
	if err != nil {
		return Purchase{}, err
	}

	return purchaseRecord, nil
}

func (u updatepurchasestatus) UpdatePurchaseStatus(ctx context.Context, purchaseRequest *UpdatePurchaseOrder, id uint) (Purchase, error) {
	status := models.PurchaseStatus(purchaseRequest.Status)

	var fromStatus models.PurchaseStatus
	updateData := models.Purchase{
		Status: status,
	}

	switch status {
	case models.Paid:
		return Purchase{}, fmt.Errorf("status paid can only be set via payment confirmation")

	case models.Received:
		fromStatus = models.OnDelivery
		now := time.Now()
		updateData.ReceivedAt = &now

	case models.Cancelled:
		return Purchase{}, fmt.Errorf("use the cancel endpoint to cancel a purchase")

	default:
		return Purchase{}, fmt.Errorf("unsupported status transition to %s", status)
	}

	result := u.repository.WithContext(ctx).Model(&models.Purchase{}).Where("id = ? AND status = ?", id, fromStatus).Updates(&updateData)
	if result.Error != nil {
		return Purchase{}, result.Error
	}
	if result.RowsAffected == 0 {
		return Purchase{}, fmt.Errorf("purchase is not a valid state for this transition, or not found")
	}

	purchaseRecord, err := u.getbypurchaseid.GetByPurchaseID(ctx, id)
	if err != nil {
		return Purchase{}, err
	}

	return purchaseRecord, nil
}
