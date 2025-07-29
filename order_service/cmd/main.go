package main

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"order/internal/domain/entity"
	"order/internal/domain/repository"
	"order/internal/domain/usecase"
	"order/internal/handler"
	"order/internal/infrastructure/postgres"
	"order/pkg/routes"
)

func main() {
	dbConn := postgres.InitDB()
	runMigrations(dbConn)

	orderRepo := repository.NewOrderRepository(dbConn)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	router := routes.SetupRouter(orderHandler)
	log.Println("Сервер запущен на :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}
