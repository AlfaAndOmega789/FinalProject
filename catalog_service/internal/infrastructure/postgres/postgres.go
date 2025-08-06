package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() *gorm.DB {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
	)

	fmt.Println("Строка подключения к БД:", connStr)

	var err error
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	//db, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatal("Ошибка подключения к БД:", err)
	//}
	//
	//if err := db.Ping(); err != nil {
	//	log.Fatal("БД недоступна:", err)
	//}
	//
	//fmt.Println("Успешное подключение к БД")
	return db
}
