package routes

import (
	"catalog/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProductByID).Methods("GET")
	router.HandleFunc("/products/", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/categories/", handlers.GetCategories).Methods("GET")

	return router
}
