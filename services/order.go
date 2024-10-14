package services

import (
	"context"
	"log"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/Marlliton/ddd-golang/domain/customer"
	custmemory "github.com/Marlliton/ddd-golang/domain/customer/memory"
	"github.com/Marlliton/ddd-golang/domain/customer/mongo"
	"github.com/Marlliton/ddd-golang/domain/product"
	prodmemory "github.com/Marlliton/ddd-golang/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customerRepo customer.CustomerRepository
	productRepo  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (
	float64, error,
) {
	// Get the customer
	c, err := os.customerRepo.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := os.productRepo.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has orderd %d products", c.GetId(), len(products))

	return total, nil
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmemory.New()

		for _, prod := range products {
			if err := pr.Add(prod); err != nil {
				return err
			}
		}

		os.productRepo = pr

		return nil
	}
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customerRepo = cr
		return nil
	}
}

func WithMemoryCostomerRepository() OrderConfiguration {
	repo := custmemory.New()

	return WithCustomerRepository(repo)
}

func WithMongoCostomerRepository(
	ctx context.Context,
	connectionString string,
) OrderConfiguration {

	return func(os *OrderService) error {
		repo, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}

		os.customerRepo = repo
		return nil
	}
}
