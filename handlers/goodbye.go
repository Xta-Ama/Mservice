package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

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
//	@Router			/goodbye [get]
func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye"))
}
