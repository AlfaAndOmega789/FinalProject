package main

import (
	"fmt"
	"log"
	"net/http"
	"order/db"
	"order/handlers"
	"order/routes"
)

func main() {
	dbConn := db.InitDB()

	//productRepo := repository.NewProductRepository(dbConn)
	//productUsecase := usecase.NewProductUsecase(productRepo)
	//productHandler := handler.NewProductHandler(productUsecase)

	router := routes.SetupRouter(productHandler)
	log.Println("Сервер запущен на :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
