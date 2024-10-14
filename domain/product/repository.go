package product

import (
	"errors"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("the product was not found")
	ErrProductAlreadyExists = errors.New("the product already exists")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(uuid.UUID) (aggregate.Product, error)
	Add(aggregate.Product) error
	Update(aggregate.Product) error
	Delete(uuid.UUID) error
}
