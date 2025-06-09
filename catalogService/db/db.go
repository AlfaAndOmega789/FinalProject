package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitDB() *sql.DB {

	_ = godotenv.Load("D:\\docker\\go_project\\catalogService\\cmd\\.env")

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

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	fmt.Println("db is nil?", db == nil)

	if err := db.Ping(); err != nil {
		log.Fatal("БД недоступна:", err)
	}

	fmt.Println("Успешное подключение к БД")
	return db
}
