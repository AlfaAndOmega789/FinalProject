package cmd

import (
	"auth/db"
	"auth/internal/user/entity"
	"auth/internal/user/handler"
	"auth/internal/user/repository"
	"auth/internal/user/usecase"
	"auth/routes"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dbConn := db.InitDB()
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
		&entity.User{},
		&entity.Role{},
		&entity.Permission{},
		&entity.RolePermission{},
	)
	if err != nil {
		log.Fatal("Migration error:", err)
	}
}
func seedRolesAndPermissions(db *gorm.DB) {
	roles := []entity.Role{
		{Name: "user", Description: "Обычный пользователь"},
		{Name: "manager", Description: "Менеджер"},
		{Name: "admin", Description: "Администратор"},
	}
	for _, r := range roles {
		db.FirstOrCreate(&r, entity.Role{Name: r.Name})
	}

	perms := []entity.Permission{
		{Code: "product.create", Description: "Создание продукта"},
		{Code: "user.manage", Description: "Управление пользователями"},
	}
	for _, p := range perms {
		db.FirstOrCreate(&p, entity.Permission{Code: p.Code})
	}
}
