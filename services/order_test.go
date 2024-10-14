package services

import (
	"testing"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	t.Helper()

	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCostomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	customer, err := aggregate.NewCustomer("Jhon")
	if err != nil {
		t.Error(err)
	}

	err = os.customerRepo.Add(customer)
	if err != nil {
		t.Error(err)
	}

	// Perform Order for one beer
	orderIDs := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(customer.GetId(), orderIDs)
	if err != nil {
		t.Error(err)
	}
}
