package services

import (
	"testing"

	"github.com/Marlliton/ddd-golang/domain/customer"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCostomerRepository(),
		// You can pass a real repository to use in the server
		// WithMongoCostomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavernService(WithOrderService(os))
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

	orderIDs := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(customer.GetId(), orderIDs)
	if err != nil {
		t.Error(err)
	}
}
