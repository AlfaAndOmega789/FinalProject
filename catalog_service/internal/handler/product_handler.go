package handler

import (
	"catalog/internal/domain/entity"
	"catalog/internal/domain/usecase"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	UseCase usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{UseCase: uc}
}

// GET /products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.UseCase.GetAll()
	if err != nil {
		http.Error(w, "Ошибка при получении продуктов", http.StatusInternalServerError)
		return
	}

	for _, p := range products {
		fmt.Fprintf(w, "ID: %s, Name: %s, Price: %.2f\n", p.ID.String(), p.Name, p.Price)
	}
}

// GET /products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}

	p, err := h.UseCase.GetByID(id)
	if err != nil {
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "ID: %s\nName: %s\nDescription: %s\nPrice: %.2f\nCategoryID: %s\nCreatedAt: %s\n",
		p.ID.String(), p.Name, p.Description, p.Price, p.CategoryID.String(), p.CreatedAt.Format("2010-01-02 10:04:05"))
}

// POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	categoryIDStr := r.FormValue("category_id")

	if name == "" || priceStr == "" || categoryIDStr == "" {
		http.Error(w, "Отсутствует имя, цена или category_id", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Неверная цена", http.StatusBadRequest)
		return
	}

	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		http.Error(w, "Неверный UUID категории", http.StatusBadRequest)
		return
	}

	product := entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}

	id, err := h.UseCase.Create(product)
	if err != nil {
		http.Error(w, "Ошибка при создании продукта", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт создан с ID: %s\n", id.String())
}

// PUT /products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	categoryIDStr := r.FormValue("category_id")

	if name == "" || priceStr == "" || categoryIDStr == "" {
		http.Error(w, "Отсутствует имя, цена или category_id", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Неверная цена", http.StatusBadRequest)
		return
	}

	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		http.Error(w, "Неверный UUID категории", http.StatusBadRequest)
		return
	}

	product := entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}

	if err := h.UseCase.Update(id, product); err != nil {
		http.Error(w, "Ошибка обновления продукта", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт с ID %s обновлён\n", id.String())
}

// DELETE /products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Неверный UUID", http.StatusBadRequest)
		return
	}

	if err := h.UseCase.Delete(id); err != nil {
		http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
