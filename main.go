package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/Xta-Ama/Mservice/docs"
	"github.com/Xta-Ama/Mservice/handlers"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

var port = os.Getenv("PORT")

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)
	r := mux.NewRouter()

	sr := r.PathPrefix("/products").Subrouter()

	sr.HandleFunc("/", ph.GetProducts).Methods(http.MethodGet)
	sr.HandleFunc("/{id:[0-9]+}", ph.DeleteProduct).Methods(http.MethodDelete)

	middlewareSr := r.PathPrefix("/products").Subrouter()
	middlewareSr.Use(ph.MiddlewareValidateProduct)

	middlewareSr.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts).Methods(http.MethodPut)
	middlewareSr.HandleFunc("/", ph.AddProduct).Methods(http.MethodPost)

	// r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
	// 	httpSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)), //The url pointing to API definition
	// 	httpSwagger.DeepLinking(true),
	// 	httpSwagger.DocExpansion("none"),
	// 	httpSwagger.DomID("swagger-ui"),
	// )).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", port),
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// go func() {
	// 	err := srv.ListenAndServe()
	// 	if err != nil {
	// 		l.Fatal(err)
	// 	}
	// }()

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// sig := <-sigChan
	// l.Println("Received terminate, graceful shutdown", sig)

	// //graceful shutdown
	// tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// srv.Shutdown(tc)

	log.Fatal(srv.ListenAndServe())
}
