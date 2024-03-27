package handlers

import (
	"net/http"

	"github.com/Xta-Ama/Mservice/api/data"
)

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

}
