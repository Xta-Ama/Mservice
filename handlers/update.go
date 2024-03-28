package handlers

import (
	"net/http"
	"strconv"

	"github.com/Xta-Ama/Mservice/models"
	"github.com/Xta-Ama/Mservice/service"
	"github.com/gorilla/mux"
)

// Update Product
//
//	@Summary		Upload file
//	@Description	Upload file
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file			true	"this is a test file"
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	string	"We need ID!!"
//	@Failure		404		{object}	string	"Can not find ID"
//	@Router			/products [put]
func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(models.Product)

	err = service.ProductUpdate(id, &prod)
	if err == service.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
