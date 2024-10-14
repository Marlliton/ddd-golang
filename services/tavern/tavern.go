package tavern

import (
	"log"

	"github.com/Marlliton/ddd-golang/services/order"
	"github.com/google/uuid"
)

type TavernConfiguration func(ts *TavernService) error

type TavernService struct {
	OrderService *order.OrderService

	// Another subservices here...
	BillingService interface{}
}

func NewTavernService(cfgs ...TavernConfiguration) (*TavernService, error) {
	ts := &TavernService{}

	for _, cfg := range cfgs {
		if err := cfg(ts); err != nil {
			return nil, err
		}
	}

	return ts, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(ts *TavernService) error {
		ts.OrderService = os
		return nil
	}
}

func (ts *TavernService) Order(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	price, err := ts.OrderService.CreateOrder(customerID, productsIDs)
	if err != nil {
		return err
	}

	log.Printf("Bil the customer %0.0f\n", price)
	return nil
}
