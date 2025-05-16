package handlers

import (
	"catalog/db"
	"catalog/models"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Conn.Query(context.TODO(), `SELECT id, name, description, price, category_id, created_at FROM products`)
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.CreatedAt)
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	row := db.Conn.QueryRow(context.Background(), "SELECT id, name, description, price, category_id, created_at FROM products WHERE id=$1", id)

	var p models.Product
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.CreatedAt)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := db.Conn.QueryRow(context.Background(),
		"INSERT INTO products (name, description, price, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		p.Name, p.Description, p.Price, p.CategoryID).Scan(&p.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Conn.Exec(context.Background(),
		"UPDATE products SET name=$1, description=$2, price=$3, category_id=$4 WHERE id=$5",
		p.Name, p.Description, p.Price, p.CategoryID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := db.Conn.Exec(context.Background(), "DELETE FROM products WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
