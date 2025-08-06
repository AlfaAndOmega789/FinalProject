package cmd

import (
	"background/internal/config"
	"background/internal/domain/entity"
	"background/internal/domain/usecase"
	"background/internal/handler"
	"background/internal/infrastructure/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	db := postgres.InitDB()

	if err := db.AutoMigrate(&entity.Currency{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	repo := postgres.NewCurrencyGormRepository(db)
	uc := usecase.NewCurrencyUseCase(repo)
	h := handler.NewCurrencyHandler(uc, cfg.CurrencyAPIURL)

	r := mux.NewRouter()
	r.HandleFunc("/rates/update", h.UpdateRates).Methods("POST")

	log.Println("Сервер запущен на :8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}
