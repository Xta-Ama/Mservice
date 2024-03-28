package data

import (
	"testing"

	models "github.com/Xta-Ama/Mservice/models"
)

func TestChecksValidation(t *testing.T) {
	p := &models.Product{
		Name:  "rice",
		Price: 1.23,
		SKU:   "abc-cdf-ifs",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
