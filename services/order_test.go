package services

import (
	"testing"

	"github.com/Marlliton/ddd-golang/domain/customer"
	"github.com/Marlliton/ddd-golang/domain/product"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []product.Product {
	t.Helper()

	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
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

	customer, err := customer.NewCustomer("Jhon")
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
