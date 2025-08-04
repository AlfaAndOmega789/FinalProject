package httpclient

import (
	"background/internal/domain/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RateResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchRates(url string) ([]entity.Currency, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить данные: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var parsed RateResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("не получилось распарсить: %w", err)
	}

	var result []entity.Currency
	for code, rate := range parsed.Rates {
		result = append(result, entity.Currency{
			Code:      code,
			Rate:      rate,
			UpdatedAt: time.Now(),
		})
	}
	return result, nil
}
