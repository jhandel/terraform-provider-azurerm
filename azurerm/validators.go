package azurerm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/satori/uuid"
)

// validateIntInSlice returns a SchemaValidateFunc which tests if the provided value
// is of type int and matches the value of an element in the valid slice
func validateIntInSlice(valid []int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(int)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be int", k))
			return
		}

		for _, str := range valid {
			if v == str {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %s", k, valid, v))
		return
	}
}

func validateUUID(v interface{}, k string) (ws []string, errors []error) {
	if _, err := uuid.FromString(v.(string)); err != nil {
		errors = append(errors, fmt.Errorf("%q is an invalid UUUID: %s", k, err))
	}
	return
}
