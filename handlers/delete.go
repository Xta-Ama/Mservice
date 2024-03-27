package handlers

import (
	"net/http"
	"strconv"

	service "github.com/Xta-Ama/Mservice/service"

	"github.com/gorilla/mux"
)

// Add product
//
// @Summary		Upload file
// @Description	Upload file
// @Accept		multipart/form-data
// @Produce		json
// @Param		file	formData	file			true	"this is a test file"
// @Success		200		{string}	string			"ok"
// @Failure		400		{object}	string	"We need ID!!"
// @Failure		404		{object}	string	"Can not find ID"
// @Router		/products [delete]
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//this will always convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE product", id)

	err := service.ProductDelete(id)

	if err == service.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
