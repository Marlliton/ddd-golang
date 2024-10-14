package memory

import (
	"fmt"
	"sync"

	"github.com/Marlliton/ddd-golang/domain/customer"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}
	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("custumer does not exist: %w", customer.ErrUdateCustumer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}
