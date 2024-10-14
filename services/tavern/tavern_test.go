package tavern

import (
	"testing"

	"github.com/Marlliton/ddd-golang/domain/product"
	"github.com/Marlliton/ddd-golang/services/order"
	"github.com/google/uuid"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOrderService(
		order.WithMemoryCostomerRepository(),
		// You can pass a real repository to use in the server
		// WithMongoCostomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavernService(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("Jhon")
	if err != nil {
		t.Error(err)
	}

	orderIDs := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(uid, orderIDs)
	if err != nil {
		t.Error(err)
	}
}

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
