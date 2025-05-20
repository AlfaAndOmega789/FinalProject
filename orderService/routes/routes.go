package routes

import (
	"github.com/gorilla/mux"
	"order/handlers"
)

func SetupRouter(productHandler *handlers.ProductHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders/", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/orders/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/orders/{id}", productHandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/orders/{id}", productHandler.PatchProduct).Methods("PATCH")
	return router
}
