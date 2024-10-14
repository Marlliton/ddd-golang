package aggregate_test

import (
	"log"
	"testing"

	"github.com/Marlliton/ddd-golang/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "shold return error if name is empty",
			name:        "",
			expectedErr: aggregate.ErrMissingValues,
		}, {
			test:        "valid values",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.description, tc.price)
			log.Println(err)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
