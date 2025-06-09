package routes

import (
	"catalog/internal/catalog/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(productHandler *handler.ProductHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	//router.HandleFunc("/categories/", productHandler.GetCategories).Methods("GET") // потом удалить(в проекте нет метода GET для категорий)
	//router.HandleFunc("/categories/{id}", productHandler.CreateCategories).Methods("PUT")
	//router.HandleFunc("/categories/{id}", productHandler.UpdateCategories).Methods("POST")
	//router.HandleFunc("/categories/{id}", productHandler.DeleteCategories).Methods("DELETE")

	return router
}
