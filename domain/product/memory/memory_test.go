package memory

import (
	"errors"
	"testing"

	"github.com/Marlliton/ddd-golang/aggregate"
	"github.com/Marlliton/ddd-golang/domain/product"
	"github.com/google/uuid"
)

func TestMemoryProduct_Add(t *testing.T) {
	repo := New()

	product, err := aggregate.NewProduct("Beer", "Good for your're health", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	repo.Add(product)

	if len(repo.products) != 1 {
		t.Errorf("expected 1 product, got %d", len(repo.products))
	}
}

func TestMemoryProduct_Get(t *testing.T) {
	repo := New()

	existingProduct, err := aggregate.NewProduct("Beer", "Good for your're health", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	repo.Add(existingProduct)
	if len(repo.products) != 1 {
		t.Errorf("expected 1 product, got %d", len(repo.products))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingProduct.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryProduct_Delete(t *testing.T) {
	repo := New()
	existingProduct, err := aggregate.NewProduct("Beer", "Good for your're health", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	repo.Add(existingProduct)
	if len(repo.products) != 1 {
		t.Errorf("expected 1 product, got %d", len(repo.products))
	}

	err = repo.Delete(existingProduct.GetID())
	if err != nil {
		t.Error(err)
	}

	size := len(repo.products)
	if size != 0 {
		t.Errorf("expected 0 products, got %d", size)
	}
}
