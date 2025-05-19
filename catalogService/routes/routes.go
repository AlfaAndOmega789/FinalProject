package routes

import (
	"catalog/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter(productHandler *handlers.ProductHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/products/", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/categories/", productHandler.GetCategories).Methods("GET")

	return router
}
