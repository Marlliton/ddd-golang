package customer

import (
	"errors"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found")
	ErrFailedToAddCustomer = errors.New("failed to add/save customer")
	ErrUdateCustumer       = errors.New("failed to updadate the custumer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
