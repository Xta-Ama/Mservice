package handlers

import (
	"net/http"
	"product-api/product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//this will always convert becasue of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
