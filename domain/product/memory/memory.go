package memory

import (
	"sync"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/Marlliton/ddd-golang/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mr.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (mr *MemoryProductRepository) Add(prod aggregate.Product) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[prod.GetID()]; !ok {
		mr.products[prod.GetID()] = prod
		return nil
	}

	return product.ErrProductAlreadyExists
}

func (mr *MemoryProductRepository) Update(prod aggregate.Product) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[prod.GetID()]; ok {
		mr.products[prod.GetID()] = prod
		return nil
	}

	return product.ErrProductNotFound
}

func (mr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[id]; ok {
		delete(mr.products, id)
		return nil
	}

	return product.ErrProductNotFound
}
