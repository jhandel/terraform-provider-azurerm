package azurerm

import "testing"

func TestValidateIntInSlice(t *testing.T) {

	cases := []struct {
		Input  []int
		Value  int
		Errors int
	}{
		{
			Input:  []int{},
			Value:  0,
			Errors: 1,
		},
		{
			Input:  []int{1},
			Value:  1,
			Errors: 0,
		},
		{
			Input:  []int{1, 2, 3, 4, 5},
			Value:  3,
			Errors: 0,
		},
		{
			Input:  []int{1, 3, 5},
			Value:  3,
			Errors: 0,
		},
		{
			Input:  []int{1, 3, 5},
			Value:  4,
			Errors: 1,
		},
	}

	for _, tc := range cases {
		_, errors := validateIntInSlice(tc.Input)(tc.Value, "azurerm_postgresql_database")

		if len(errors) != tc.Errors {
			t.Fatalf("Expected the validateIntInSlice trigger a validation error for input: %+v looking for %+v", tc.Input, tc.Value)
		}
	}

}
