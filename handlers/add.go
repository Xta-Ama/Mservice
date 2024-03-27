package handlers

import (
	"net/http"

	models "github.com/Xta-Ama/Mservice/models"
	service "github.com/Xta-Ama/Mservice/service"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(models.Product)
	service.ProductAdd(&prod)
}
