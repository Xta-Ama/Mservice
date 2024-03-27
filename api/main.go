package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Xta-Ama/Mservice/api/handlers"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

var port = os.Getenv("PORT")

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// creating a new serve mux
	r := mux.NewRouter()

	sr := r.PathPrefix("/products").Subrouter()
	sr.HandleFunc("/", ph.GetProducts).Methods(http.MethodGet)
	sr.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct).Methods(http.MethodDelete)

	middlewareSr := r.PathPrefix("/products").Subrouter()
	middlewareSr.Use(ph.MiddlewareValidateProduct)
	middlewareSr.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts).Methods(http.MethodPut)
	middlewareSr.HandleFunc("/", ph.AddProduct).Methods(http.MethodPost)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	//graceful shutdown
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(tc)

	// log.Fatal(srv.ListenAndServe())
}
