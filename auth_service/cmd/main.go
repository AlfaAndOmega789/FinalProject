package cmd

import (
	entity2 "auth/internal/domain/entity"
	"auth/internal/domain/repository"
	"auth/internal/domain/usecase"
	"auth/internal/handler"
	"auth/internal/infrastructure/postgres"
	"auth/pkg/routes"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dbConn := postgres.InitDB()
	runMigrations(dbConn)
	seedRolesAndPermissions(dbConn)

	authRepo := repository.NewAuthRepository(dbConn)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	router := routes.SetupRouter(authHandler)
	log.Println("Сервер запущен на :8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity2.User{},
		&entity2.Role{},
		&entity2.Permission{},
		&entity2.RolePermission{},
	)
	if err != nil {
		log.Fatal("Migration error:", err)
	}
}
func seedRolesAndPermissions(db *gorm.DB) {
	roles := []entity2.Role{
		{Name: "user", Description: "Обычный пользователь"},
		{Name: "manager", Description: "Менеджер"},
		{Name: "admin", Description: "Администратор"},
	}
	for _, r := range roles {
		db.FirstOrCreate(&r, entity2.Role{Name: r.Name})
	}

	perms := []entity2.Permission{
		{Code: "product.create", Description: "Создание продукта"},
		{Code: "user.manage", Description: "Управление пользователями"},
	}
	for _, p := range perms {
		db.FirstOrCreate(&p, entity2.Permission{Code: p.Code})
	}
}
