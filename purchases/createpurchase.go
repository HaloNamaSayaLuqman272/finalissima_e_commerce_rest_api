package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"finalissima_e_commerce_rest_api/package/constant"
	"finalissima_e_commerce_rest_api/package/rajaongkir"
	"finalissima_e_commerce_rest_api/package/utils"
	"finalissima_e_commerce_rest_api/products"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type createpurchase struct {
	repository         *gorm.DB
	productService     products.Service
	roService          rajaongkir.Service
	shippingService    ShippingService
	getbypurchaseid    getbypurchaseid
	shopOriginDistrict string
}

func NewCreatePurchase(repository *gorm.DB, productService products.Service, roService rajaongkir.Service, shippingService ShippingService, getPurchase getbypurchaseid) createpurchase {
	return createpurchase{
		repository:         repository,
		productService:     productService,
		roService:          roService,
		shippingService:    shippingService,
		getbypurchaseid:    getPurchase,
		shopOriginDistrict: utils.GetConfigurance(constant.SHOP_ORIGIN_DISTRICT_ID),
	}
}

func (c createpurchase) CreatePurchase(ctx context.Context, purchaseOrder *PurchaseOrder) (Purchase, error) {
	var totalWeight uint
	var totalAmount float64
	items := make([]models.PurchaseItem, 0, len(purchaseOrder.Items))

	for _, itemPurchaseOrder := range purchaseOrder.Items {
		product, err := c.productService.GetProductByID(ctx, itemPurchaseOrder.ProductID)
		if err != nil {
			return Purchase{}, fmt.Errorf("product id %d not found", itemPurchaseOrder.ProductID)
		}

		if product.Stock < itemPurchaseOrder.Quantity {
			return Purchase{}, fmt.Errorf("product stock %s not enough", product.NameProduct)
		}

		totalWeight += product.Weight * itemPurchaseOrder.Quantity
		totalAmount += product.Price * float64(itemPurchaseOrder.Quantity)

		items = append(items, models.PurchaseItem{
			ProductID: itemPurchaseOrder.ProductID,
			Quantity:  itemPurchaseOrder.Quantity,
			Weight:    product.Weight,
			Price:     product.Price,
		})
	}

	purchaseOrder.TotalWeight = totalWeight
	purchaseOrder.Amount = totalAmount

	fee, etd, err := c.roService.GetDeliveryFee(rajaongkir.GetFeeRequest{
		Origin:      c.shopOriginDistrict,
		Destination: fmt.Sprintf("%d", purchaseOrder.DestinationID),
		Weight:      totalWeight,
		Courier:     purchaseOrder.Courier,
	})
	if err != nil {
		return Purchase{}, fmt.Errorf("failed calculate shipping cost: %w", err)
	}
	purchaseOrder.Fee = fee

	days, ok := ParseETDDays(etd)
	if !ok {
		days = DefaultEstimatedDays
	}
	purchaseOrder.EstimatedDelivery = days
	estimated := time.Now().AddDate(0, 0, int(days))
	purchaseOrder.EstimatedReceivedAt = &estimated

	purchase := models.Purchase{
		UserID:              purchaseOrder.UserID,
		Items:               items,
		TotalWeight:         purchaseOrder.TotalWeight,
		Amount:              purchaseOrder.Amount,
		DestinationID:       purchaseOrder.DestinationID,
		Fee:                 purchaseOrder.Fee,
		Courier:             purchaseOrder.Courier,
		Status:              models.Pending,
		EstimatedDelivery:   purchaseOrder.EstimatedDelivery,
		EstimatedReceivedAt: purchaseOrder.EstimatedReceivedAt,
	}

	err = c.repository.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&purchase).Error; err != nil {
			return err
		}

		for _, item := range items {
			if err := tx.Model(&models.Product{}).Where("id = ?", item.ProductID).Update("stock", gorm.Expr("stock - ?", item.Quantity)).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return Purchase{}, err
	}

	return c.getbypurchaseid.GetByPurchaseID(ctx, purchase.ID)
}
