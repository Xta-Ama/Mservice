package handlers

import (
	"net/http"
	"product-api/product-api/data"
)

// swagger:route GET /products products listproducts
// Returns a list of product
// reponses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Hnadle GET Products")

	//fetch the products from hte datsstore
	lp := data.GetProducts()

	//serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
