package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "guri",
		Price: 2,
		SKU:   "abs-fds-dcv",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}

}
