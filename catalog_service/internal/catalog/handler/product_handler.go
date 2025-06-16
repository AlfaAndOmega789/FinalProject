package handler

import (
	"catalog/internal/catalog/entity"
	"catalog/internal/catalog/usecase"
	"fmt"
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
		fmt.Fprintf(w, "ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

// GET /products/{id}
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный продукт ID", http.StatusBadRequest)
		return
	}

	p, err := h.UseCase.GetByID(id)
	if err != nil {
		http.Error(w, "Продукт не найден", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "ID: %d\nName: %s\nDescription: %s\nPrice: %.2f\nCategoryID: %v\nCreatedAt: %s\n",
		p.ID, p.Name, p.Description, p.Price, p.CategoryID, p.CreatedAt.Format("2006-01-02 15:04:05"))
}

// POST /products
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	categoryIDStr := r.FormValue("category_id")

	if name == "" || priceStr == "" {
		http.Error(w, "Отсутствует имя или цена", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Неверная цена", http.StatusBadRequest)
		return
	}

	var categoryID int
	if categoryIDStr != "" {
		catID, err := strconv.ParseInt(categoryIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Неверная category_id", http.StatusBadRequest)
			return
		}
		categoryID = int(catID)
	}

	product := entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  uint(categoryID),
	}

	id, err := h.UseCase.Create(product)
	if err != nil {
		http.Error(w, "Ошибка при создании продукта", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт создан ID: %d\n", id)
}

// PUT /products/{id}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	categoryIDStr := r.FormValue("category_id")

	if name == "" || priceStr == "" {
		http.Error(w, "Отсутствует имя или цена", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Неверная цена", http.StatusBadRequest)
		return
	}

	var categoryID int
	if categoryIDStr != "" {
		catID, err := strconv.ParseInt(categoryIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Неверный category_id", http.StatusBadRequest)
			return
		}
		categoryID = int(catID)
	}

	product := entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  uint(categoryID),
	}

	err = h.UseCase.Update(id, product)
	if err != nil {
		http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт с ID %d обновлен\n", id)
}

// DELETE /products/{id}
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "Продукт с ID %d удален\n", id)
}
