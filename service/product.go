package service

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	models "github.com/Xta-Ama/Mservice/models"
)

type Products []*models.Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func ProductGetAll() Products {
	return productList
}

func ProductAdd(p *models.Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func ProductDelete(id int) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

func ProductUpdate(id int, p *models.Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*models.Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*models.Product{
	{
		ID:          1,
		Name:        "Milo",
		Description: "Best for breakfast",
		Price:       2.24,
		SKU:         "abc1223",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Lipton",
		Description: "I like the smell",
		Price:       3.34,
		SKU:         "ghs234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
