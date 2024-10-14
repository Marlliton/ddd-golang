package memory

import (
	"sync"

	"github.com/Marlliton/ddd-golang/domain/product"
	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product

	for _, product := range mr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mr.products[id]; ok {
		return product, nil
	}

	return product.Product{}, product.ErrProductNotFound
}

func (mr *MemoryProductRepository) Add(prod product.Product) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[prod.GetID()]; !ok {
		mr.products[prod.GetID()] = prod
		return nil
	}

	return product.ErrProductAlreadyExists
}

func (mr *MemoryProductRepository) Update(prod product.Product) error {
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
