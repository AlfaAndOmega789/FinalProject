package main

import (
	"catalog/internal/domain/entity"
	"catalog/internal/domain/repository"
	"catalog/internal/domain/usecase"
	handler2 "catalog/internal/handler"
	"catalog/internal/infrastructure/postgres"
	"catalog/pkg/routes"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dbConn := postgres.InitDB()
	runMigrations(dbConn)

	productRepo := repository.NewProductRepository(dbConn)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler2.NewProductHandler(productUsecase)

	categoryRepo := repository.NewCategoryRepository(dbConn)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler2.NewCategoryHandler(categoryUsecase)

	router := routes.SetupRouter(productHandler, categoryHandler)
	log.Println("Сервер запущен на :8081")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Category{}, &entity.Product{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}
