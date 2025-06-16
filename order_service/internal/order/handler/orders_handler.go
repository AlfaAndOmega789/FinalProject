package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"order/internal/order/entity"
	"order/internal/order/usecase"
	"strconv"
)

type OrderHandler struct {
	usecase *usecase.OrderUsecase
}

func NewOrderHandler(u *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: u}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	order := &entity.Order{}
	order.UserID = r.FormValue("user_id")
	order.Currency = r.FormValue("currency")
	order.Status = r.FormValue("status")

	total, err := strconv.ParseFloat(r.FormValue("total_price"), 64)
	if err != nil {
		http.Error(w, "неверная total_price", http.StatusBadRequest)
		return
	}
	order.TotalPrice = total

	delivery, err := strconv.ParseFloat(r.FormValue("delivery_price"), 64)
	if err != nil {
		http.Error(w, "неверная delivery_price", http.StatusBadRequest)
		return
	}
	order.DeliveryPrice = delivery

	if err := h.usecase.Create(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Заказ создан: %s", order.ID)
}

func (h *OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := fmt.Sprintf("Заказ: ID=%s, Статус=%s, Общая стоимость=%.2f %s\n", order.ID, order.Status, order.TotalPrice, order.Currency)
	fmt.Fprint(w, response)
}

func (h *OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.usecase.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *OrderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	status := r.FormValue("status")

	if status == "" {
		http.Error(w, "статус отсутствует", http.StatusBadRequest)
		return
	}

	if err := h.usecase.Update(id, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Заказ %s обновленный статус: %s\n", id, status)
}
