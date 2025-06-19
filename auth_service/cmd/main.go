package cmd

import (
	"auth/internal/user/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
	"reviews/routes"
)

func main() {
	dbConn := db.InitDB()
	runMigrations(dbConn)

	authRepo := repository.NewAuthRepository(dbConn)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	router := routes.SetupRouter(authHandler)
	log.Println("Сервер запущен на :8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Auth{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}
