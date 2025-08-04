package handler

import (
	"background/internal/domain/usecase"
	"background/pkg/httpclient"
	"net/http"
)

type CurrencyHandler struct {
	uc     *usecase.CurrencyUseCase
	apiURL string
}

func NewCurrencyHandler(uc *usecase.CurrencyUseCase, url string) *CurrencyHandler {
	return &CurrencyHandler{
		uc:     uc,
		apiURL: url,
	}
}

func (h *CurrencyHandler) UpdateRates(w http.ResponseWriter, r *http.Request) {
	rates, err := httpclient.FetchRates(h.apiURL)
	if err != nil {
		http.Error(w, "ошибка получения из данных их API: "+err.Error(), http.StatusBadGateway)
		return
	}

	if err := h.uc.SaveRates(rates); err != nil {
		http.Error(w, "ошибка сохранени в базе : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("курсы успешно обновлены"))
}
