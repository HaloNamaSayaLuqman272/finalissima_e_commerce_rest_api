package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"log"

	"gorm.io/gorm"
)

type createpurchase struct {
	repository *gorm.DB
}

func (c createpurchase) CreatePurchase(ctx context.Context, purchaseOrder *PurchaseOrder) error {
	err := c.repository.Transaction(func(tx *gorm.DB) error {
		// aplikasi Go melakukan cek terhadap semua product yg ada di database
		for _, productId := range purchaseOrder.Products {
			var product models.Product
			if err := tx.WithContext(ctx).Where("id IN ? AND stock > 0", productId).Find(&product).Error; err != nil {
				log.Println("DEBUG error at:", err)
				return err
			}

			purchase := models.Purchase{
				UserID:    purchaseOrder.UserID,
				ProductID: product.ID,
				Fee:       purchaseOrder.Fee,
				Courier:   purchaseOrder.Courier,
				Status:    models.Pending,
			}
			if err := tx.WithContext(ctx).Create(&purchase).Error; err != nil {
				log.Printf("DEBUG error at: %v", err)
				return err
			}

			if err := tx.WithContext(ctx).First(&product, "id = ?", productId).Error; err != nil {
				log.Printf("DEBUG error at: %v", err)
				return err
			}
			product.Stock = product.Stock - 1
			if err := tx.WithContext(ctx).Where("id = ?", productId).Updates(product).Error; err != nil {
				log.Printf("DEBUG error at: %v", err)
				return err
			}
		}

		return nil
	})

	return err
}
