package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found")
	ErrFailedToAddCustomer = errors.New("failed to add/save customer")
	ErrUdateCustumer       = errors.New("failed to updadate the custumer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}
