package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/database/models"
	"fmt"

	"gorm.io/gorm"
)

type deletepurchase struct {
	repository      *gorm.DB
	getbypurchaseid getbypurchaseid
}

func NewDeletePurchase(repository *gorm.DB, getbypurchaseid getbypurchaseid) deletepurchase {
	return deletepurchase{
		repository:      repository,
		getbypurchaseid: getbypurchaseid,
	}
}

func (d deletepurchase) DeletePurchase(ctx context.Context, id uint) error {
	purchase, err := d.getbypurchaseid.GetByPurchaseID(ctx, id)
	if err != nil {
		return err
	}

	if purchase.Status == models.OnDelivery || purchase.Status == models.Received || purchase.Status == models.Cancelled {
		return fmt.Errorf("deleted purchase not available in purchase status: %s", purchase.Status)
	}

	if err := d.repository.WithContext(ctx).Delete(&models.Purchase{}).Error; err != nil {
		return err
	}

	return nil
}
