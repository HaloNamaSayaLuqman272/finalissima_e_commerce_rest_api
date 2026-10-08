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

func (c createpurchase) CreatePurchase(ctx context.Context, purchaseOrder *PurchaseOrder) error {
	// err := c.repository.Transaction(func(tx *gorm.DB) error {
	// 	var totalWeight uint
	// 	var totalAmount float64
	// 	items := make([]models.PurchaseItem, 0, len(purchaseOrder.Items))

	// 	for _, productId := range purchaseOrder.Items {

	// 		// check if product exists and has stock
	// 		var product models.Product
	// 		if err := tx.WithContext(ctx).Where("id = ? AND stock > 0", productId).Find(&product).Error; err != nil {
	// 			return err
	// 		}

	// 		// make purchase order
	// 		totalWeight += product.Weight * productId.Quantity
	// 		totalAmount += product.Price * float64(productId.Quantity)

	// 		items = append(items, models.PurchaseItem{
	// 			ProductID: productId.ProductID,
	// 			Quantity:  productId.Quantity,
	// 			Weight:    product.Weight,
	// 			Price:     product.Price,
	// 		})

	// 		purchaseOrder.TotalWeight = totalWeight
	// 		purchaseOrder.Amount = totalAmount

	// 		fee, etd, err := c.roService.GetDeliveryFee(ctx, rajaongkir.GetFeeRequest{
	// 			Origin:      c.shopOriginDistrict,
	// 			Destination: fmt.Sprintf("%d", purchaseOrder.DestinationID),
	// 			Weight:      totalWeight,
	// 			Courier:     purchaseOrder.Courier,
	// 		})
	// 		if err != nil {
	// 			return err
	// 		}

	// 		purchaseOrder.Fee = fee

	// 		days, ok := ParseETDDays(etd)
	// 		if !ok {
	// 			days = DefaultEstimatedDays
	// 		}
	// 		purchaseOrder.EstimatedDelivery = days
	// 		estimated := time.Now().AddDate(0, 0, int(days))
	// 		purchaseOrder.EstimatedReceivedAt = &estimated

	// 		// create purchase order
	// 		purchase := models.Purchase{
	// 			UserID:              purchaseOrder.UserID,
	// 			Items:               items,
	// 			TotalWeight:         purchaseOrder.TotalWeight,
	// 			Amount:              purchaseOrder.Amount,
	// 			DestinationID:       purchaseOrder.DestinationID,
	// 			Fee:                 purchaseOrder.Fee,
	// 			Courier:             purchaseOrder.Courier,
	// 			Status:              models.Pending,
	// 			EstimatedDelivery:   purchaseOrder.EstimatedDelivery,
	// 			EstimatedReceivedAt: purchaseOrder.EstimatedReceivedAt,
	// 		}

	// 		// save purchase order to database
	// 		if err := tx.WithContext(ctx).Create(&purchase).Error; err != nil {
	// 			return err
	// 		}

	// 		// update product stock
	// 		if err := tx.WithContext(ctx).Model(&models.Product{}).Where("id = ?", productId.ProductID).Update("stock", gorm.Expr("stock - ?", productId.Quantity)).Error; err != nil {
	// 			return err
	// 		}
	// 	}

	// 	return nil
	// })

	// return err
	var totalWeight uint
	var totalAmount float64
	items := make([]models.PurchaseItem, 0, len(purchaseOrder.Items))

	for _, itemPurchaseOrder := range purchaseOrder.Items {
		product, err := c.productService.GetProductByID(ctx, itemPurchaseOrder.ProductID)
		if err != nil {
			return err
		}

		if product.Stock < itemPurchaseOrder.Quantity {
			return err
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

	fee, etd, err := c.roService.GetDeliveryFee(ctx, rajaongkir.GetFeeRequest{
		Origin:      c.shopOriginDistrict,
		Destination: fmt.Sprintf("%d", purchaseOrder.DestinationID),
		Weight:      totalWeight,
		Courier:     purchaseOrder.Courier,
	})
	if err != nil {
		return err
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
	// if err != nil {
	// 	return Purchase{}, err
	// }

	// fmt.Println("DEBUG purchase.ID setelah transaksi =", purchase.ID)

	// record := new(Purchase)
	// fmt.Println("DEBUG record sebelum First tanpa preload")
	// if err := c.repository.WithContext(ctx).First(record, purchase.ID).Error; err != nil {
	// 	fmt.Println("DEBUG First tanpa preload error")
	// 	return Purchase{}, err
	// }
	// fmt.Println("DEBUG sesudah First tanpa preload =", record.ID)

	// if err := c.repository.WithContext(ctx).Model(record).Association("Items").Find(&record.Items); err != nil {
	// 	fmt.Println("DEBUG association items error:", err)
	// 	return Purchase{}, err
	// }

	// fmt.Println("DEBUG record setelah load items, jumlah items =", len(record.Items))

	// return *record, err
	return err
}
