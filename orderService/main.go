package main

import (
	"fmt"
	"log"
	"net/http"
	"order/db"
	"order/handlers"
	"order/routes"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Order service is working!")
}

func main() {
	database := db.InitDB()
	productHandler := &handlers.ProductHandler{DB: database}

	router := routes.SetupRouter(productHandler)

	log.Println("Сервер запущен на :8082")
	http.ListenAndServe(":8082", router)
}
