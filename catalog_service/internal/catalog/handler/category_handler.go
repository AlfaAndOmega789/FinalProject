package handler

import (
	"catalog/internal/catalog/entity"
	"catalog/internal/catalog/usecase"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	UseCase usecase.CategoryUsecase
}

func NewCategoryHandler(uc usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{UseCase: uc}
}

// GET /categories
func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := h.UseCase.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении категорий", http.StatusInternalServerError)
		return
	}

	for _, p := range categories {
		fmt.Fprintf(w, "ID: %d, Name: %s\n", p.ID, p.Name)
	}
}

// POST /categories
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Отсутствует название категории", http.StatusBadRequest)
		return
	}

	category := entity.Category{
		Name: name,
	}

	id, err := h.UseCase.Create(category)
	if err != nil {
		http.Error(w, "Ошибка при создании категории", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Категория создана ID: %d\n", id)
}

// PUT /categories/{id}
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	if name == "" {
		http.Error(w, "Отсутствует имя категории", http.StatusBadRequest)
		return
	}

	category := entity.Category{
		Name: name,
	}

	err = h.UseCase.Update(id, category)
	if err != nil {
		http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Категория с ID %d обновлена\n", id)
}

// DELETE /categories/{id}
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	err = h.UseCase.Delete(id)
	if err != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Категория с ID %d удален\n", id)
}
