package handlers

import (
	"net/http"
	"strconv"

	service "github.com/Xta-Ama/Mservice/service"

	"github.com/gorilla/mux"
)

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
