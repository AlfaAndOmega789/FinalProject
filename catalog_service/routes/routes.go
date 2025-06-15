package routes

import (
	"catalog/internal/catalog/handler"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter(productHandler *handler.ProductHandler, categoryHandler *handler.CategoryHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductByID).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/categories", categoryHandler.GetCategory).Methods("GET")
	router.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	router.Walk(func(route *mux.Route, r *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Printf("Зарегистрирован маршрут: %s %v\n", pathTemplate, methods)
		return nil
	})

	return router
}
