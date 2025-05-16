package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"catalog/handlers"
)

func main() {
	r := mux.NewRouter()

	// Эндпоинты
	r.HandleFunc("/products", handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	log.Println("Catalog service running on port 8081")
	http.ListenAndServe(":8081", r)
}
