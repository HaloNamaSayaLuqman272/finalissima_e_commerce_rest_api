package purchases

import (
	"context"
	"fmt"
	"time"
)

type ShippingCostResult struct {
	Fee float64
	ETD string
}

type ShippingRequest struct {
	OrderID       uint
	DestinationID uint
	Courier       string
	Weight        uint
}

type ShippingResult struct {
	TrackingNumber string
}

type ShippingService interface {
	CalculateCost(ctx context.Context, destinationID uint, weight uint, courier string) (*ShippingCostResult, error)
	CreateShipmentOrder(ctx context.Context, req ShippingRequest) (*ShippingResult, error)
	CancelShipmentOrder(ctx context.Context, trackingNumber string) error
}

type FakeShippingService struct{}

func NewFakeShippingService() FakeShippingService {
	return FakeShippingService{}
}

func (f FakeShippingService) CalculateCost(ctx context.Context, destinationID uint, weight uint, courier string) (*ShippingCostResult, error) {
	fee := float64(weight) / 1000 * 5000
	if fee < 10000 {
		fee = 10000
	}
	return &ShippingCostResult{
		Fee: fee,
		ETD: "2-3 day",
	}, nil
}

func (f FakeShippingService) CreateShipmentOrder(ctx context.Context, req ShippingRequest) (*ShippingResult, error) {
	return &ShippingResult{
		TrackingNumber: fmt.Sprintf("FAKE-RESI-%d-%d", req.OrderID, time.Now().Unix()),
	}, nil
}

func (f FakeShippingService) CancelShipmentOrder(ctx context.Context, trackingNumber string) error {
	return nil
}

type FakeShippingServiceFail struct{}

func (f FakeShippingServiceFail) CalculateCost(ctx context.Context, destinationID uint, weight uint, courier string) (*ShippingCostResult, error) {
	return nil, fmt.Errorf("destination not found or the courier doesn't serve this area")
}

func (f FakeShippingServiceFail) CreateShipmentOrder(ctx context.Context, req ShippingRequest) (*ShippingResult, error) {
	return nil, fmt.Errorf("destination not found or the courier doesn't serve this area")
}

func (f FakeShippingServiceFail) CancelShipmentOrder(ctx context.Context, trackingNumber string) error {
	return fmt.Errorf("failed to cancel shipment at Komerce for tracking number %s", trackingNumber)
}
