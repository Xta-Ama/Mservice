package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "rice",
		Price: 1.23,
		SKU:   "abc-cdf-ifs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
