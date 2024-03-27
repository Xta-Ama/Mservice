package handlers

import (
	"net/http"
	"strconv"

	"github.com/Xta-Ama/Mservice/models"
	"github.com/Xta-Ama/Mservice/service"
	"github.com/gorilla/mux"
)

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
