package memory

import (
	"errors"
	"testing"

	"github.com/Marlliton/ddd-golang/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := customer.NewCustomer("Jhon")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetId()
	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.New(),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add a customer",
			cust:        "Jhon",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]customer.Customer{},
			}

			cust, err := customer.NewCustomer("Jhon")
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("expeted error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetId())
			if err != nil {
				t.Fatal(err)
			}

			if found.GetId() != cust.GetId() {
				t.Errorf("expeted %v, got %v", cust.GetId(), found.GetId())
			}
		})
	}

}
