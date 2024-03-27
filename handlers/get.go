package handlers

import (
	"net/http"

	service "github.com/Xta-Ama/Mservice/service"
)

// GetProducts example
//
//	@Summary		Get all the products from the store
//	@Description	get all products
//	@Accept			json
//	@Produce		json
//	@Param			some_id	path		int				true	"Some ID"
//	@Param			some_id	body		web.Pet			true	"Some ID"
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	web.APIError	"Something went wrong"
//	@Failure		404		{object}	web.APIError	"Something went wrong"
//	@Router			/products [get]
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := service.ProductGetAll()

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
