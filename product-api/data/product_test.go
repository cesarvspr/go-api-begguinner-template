package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "guri",
		Price: 2,
		SKU:   "abs-fds-dcv",
	}
	v := NewValidation()
	err := v.Validate(p)
	if err != nil {
		t.Fatal(err)
	}

}
