package handlers

import (
	"net/http"

	models "github.com/Xta-Ama/Mservice/models"
	service "github.com/Xta-Ama/Mservice/service"
)

// Add product
//
//	@Summary		Upload file
//	@Description	Upload file
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file			true	"this is a test file"
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{object}	string	"We need ID!!"
//	@Failure		404		{object}	string	"Can not find ID"
//	@Router			/products [post]
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(models.Product)
	service.ProductAdd(&prod)
}
