package purchases

import (
	"context"
	"finalissima_e_commerce_rest_api/package/rajaongkir"
	"finalissima_e_commerce_rest_api/package/utils"
	"finalissima_e_commerce_rest_api/products"

	"gorm.io/gorm"
)

type Service interface {
	CreatePurchase(ctx context.Context, purchaseOrder *PurchaseOrder) error
	GetByPurchaseID(ctx context.Context, id uint) (Purchase, error)
	GetPurchaseByUserID(ctx context.Context, userId uint) ([]Purchase, error)
	GetAllPurchases(ctx context.Context, pagination utils.Pagination) ([]Purchase, error)
	ConfirmPayment(ctx context.Context, id uint) (Purchase, error)
	UpdatePurchaseStatus(ctx context.Context, purchaseRequest *UpdatePurchaseOrder, id uint) (Purchase, error)
	CancelPurchase(ctx context.Context, id uint, requestingUserID uint, isAdmin bool) (Purchase, error)
	DeletePurchase(ctx context.Context, id uint) error
}

type service struct {
	createpurchase
	getbypurchaseid
	getpurchasesbyuserid
	getallpurchases
	updatepurchasestatus
	cancelpurchase
	deletepurchase
}

var _ Service = (*service)(nil)

func New(repository *gorm.DB, productService products.Service, roService rajaongkir.Service, shippingService ShippingService) Service {
	getPurchase := getbypurchaseid{repository: repository}

	return service{
		createpurchase:       NewCreatePurchase(repository, productService, roService, shippingService, getPurchase),
		getbypurchaseid:      getPurchase,
		getpurchasesbyuserid: getpurchasesbyuserid{repository: repository},
		getallpurchases:      getallpurchases{repository: repository},
		updatepurchasestatus: NewUpdatePurchaseStatus(repository, getPurchase, shippingService),
		cancelpurchase:       NewCancelPurchase(repository, getPurchase, shippingService),
		deletepurchase:       NewDeletePurchase(repository, getPurchase),
	}
}

func (s service) CreatePurchase(ctx context.Context, purchaseOrder *PurchaseOrder) error {
	return s.createpurchase.CreatePurchase(ctx, purchaseOrder)
}

func (s service) GetByPurchaseID(ctx context.Context, id uint) (Purchase, error) {
	return s.getbypurchaseid.GetByPurchaseID(ctx, id)
}

func (s service) GetPurchaseByUserID(ctx context.Context, userId uint) ([]Purchase, error) {
	return s.getpurchasesbyuserid.GetPurchasesByUserID(ctx, userId)
}

func (s service) GetAllPurchases(ctx context.Context, pagination utils.Pagination) ([]Purchase, error) {
	return s.getallpurchases.GetAllPurchases(ctx, pagination)
}

func (s service) ConfirmPayment(ctx context.Context, id uint) (Purchase, error) {
	return s.updatepurchasestatus.ConfirmPayment(ctx, id)
}

func (s service) UpdatePurchaseStatus(ctx context.Context, purchaseRequest *UpdatePurchaseOrder, id uint) (Purchase, error) {
	return s.updatepurchasestatus.UpdatePurchaseStatus(ctx, purchaseRequest, id)
}

func (s service) CancelPurchase(ctx context.Context, id uint, requestingUserID uint, isAdmin bool) (Purchase, error) {
	return s.cancelpurchase.CancelPurchase(ctx, id, requestingUserID, isAdmin)
}

func (s service) DeletePurchase(ctx context.Context, id uint) error {
	return s.deletepurchase.DeletePurchase(ctx, id)
}
