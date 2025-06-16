package main

import (
	"catalog/db"
	"catalog/internal/catalog/entity"
	"catalog/internal/catalog/handler"
	"catalog/internal/catalog/repository"
	"catalog/internal/catalog/usecase"
	"catalog/routes"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dbConn := db.InitDB()
	runMigrations(dbConn)

	productRepo := repository.NewProductRepository(dbConn)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	categoryRepo := repository.NewCategoryRepository(dbConn)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

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
