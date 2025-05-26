package handlers

import (
	"catalog/models"
	//"FinalProject/db"
	//"FinalProject/models"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	DB *sql.DB
}

// GET /products
func (db *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var rows, err = db.DB.Query("SELECT id, name, description, price, category_id, created_at FROM products")
	if err != nil {
		http.Error(w, "Ошибка запроса DB", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.CreatedAt)
		if err != nil {
			continue
		}
		fmt.Fprintf(w, "ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

// GET /products/{id}
func (db *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный продукт ID", http.StatusBadRequest)
		return
	}

	var p models.Product
	err = db.DB.QueryRow("SELECT id, name, description, price, category_id, created_at FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Продукт не найден", http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка запроса DB", http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintf(w, "ID: %d\nName: %s\nDescription: %s\nPrice: %.2f\nCategoryID: %v\nCreatedAt: %s\n",
		p.ID, p.Name, p.Description.String, p.Price, p.CategoryID.Int64, p.CreatedAt.Format("2006-01-02 15:04:05"))
}

// POST /products
func (db *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
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

	var categoryID sql.NullInt64
	if categoryIDStr != "" {
		catID, err := strconv.ParseInt(categoryIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Неверная category_id", http.StatusBadRequest)
			return
		}
		categoryID = sql.NullInt64{Int64: catID, Valid: true}
	} else {
		categoryID = sql.NullInt64{Valid: false}
	}

	query := `INSERT INTO products (name, description, price, category_id) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err = db.DB.QueryRow(query, name, description, price, categoryID).Scan(&id)
	if err != nil {
		http.Error(w, "Ошибка при вставке продукта", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт создан ID: %d\n", id)
}

// PUT /products/{id}
func (db *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный продукт ID", http.StatusBadRequest)
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
		http.Error(w, "Ошибка в цене", http.StatusBadRequest)
		return
	}

	var categoryID sql.NullInt64
	if categoryIDStr != "" {
		catID, err := strconv.ParseInt(categoryIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Неверный category_id", http.StatusBadRequest)
			return
		}
		categoryID = sql.NullInt64{Int64: catID, Valid: true}
	} else {
		categoryID = sql.NullInt64{Valid: false}
	}

	query := `UPDATE products SET name = $1, description = $2, price = $3, category_id = $4 WHERE id = $5`
	_, err = db.DB.Exec(query, name, description, price, categoryID, id)
	if err != nil {
		http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт с ID %d обновлен\n", id)
}

// DELETE /products/{id}
func (db *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Нет продукта с таким ID", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Ошибка, продукт не удалился", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Продукт с ID %d удален\n", id)
}

// GET /categories
func (db *ProductHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, created_at FROM categories")
	if err != nil {
		http.Error(w, "Ошибка запроса категорий", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt)
		if err != nil {
			continue
		}
		fmt.Fprintf(w, "ID: %s, Name: %s\n", c.ID, c.Name)
	}
}
