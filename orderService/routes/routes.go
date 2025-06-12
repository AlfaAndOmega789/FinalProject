package routes

import (
	"github.com/gorilla/mux"
	"order/internal/order/handler"
)

func SetupRouter(productHandler *handler.ProductHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders/", productHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", productHandler.GetOrderByID).Methods("GET")
	router.HandleFunc("/orders/{id}", productHandler.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/orders/{id}", productHandler.PatchOrder).Methods("PATCH")
	return router
}
