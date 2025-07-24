package handler

import (
	"catalog/internal/domain/entity"
	"catalog/internal/domain/usecase"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type CategoryHandler struct {
	Usecase usecase.CategoryUsecase
}

func NewCategoryHandler(u usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{Usecase: u}
}

// GET /categories
func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Usecase.GetAll()
	if err != nil {
		http.Error(w, "Ошибка получения категорий", http.StatusInternalServerError)
		return
	}

	for _, c := range categories {
		fmt.Fprintf(w, "ID: %s, Name: %s\n", c.ID.String(), c.Name)
	}
}

// POST /categories
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Имя категории обязательно", http.StatusBadRequest)
		return
	}

	category := entity.Category{
		Name: name,
	}

	id, err := h.Usecase.Create(category)
	if err != nil {
		http.Error(w, "Ошибка при создании категории", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Категория создана с ID: %s\n", id.String())
}

// PUT /categories/{id}
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Имя категории обязательно", http.StatusBadRequest)
		return
	}

	category := entity.Category{
		Name: name,
	}

	if err := h.Usecase.Update(id, category); err != nil {
		http.Error(w, "Ошибка при обновлении категории", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Категория с ID %s обновлена\n", id.String())
}

// DELETE /categories/{id}
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}

	if err := h.Usecase.Delete(id); err != nil {
		http.Error(w, "Ошибка при удалении категории", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
