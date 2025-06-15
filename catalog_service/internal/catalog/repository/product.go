package repository

import (
	"catalog/internal/catalog/entity"
	"database/sql"
	"fmt"
)

type ProductPostgresRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{DB: db}
}

func (r *ProductPostgresRepository) GetAll() ([]entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, description, price, category_id, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductPostgresRepository) GetByID(id int) (entity.Product, error) {
	var p entity.Product
	query := "SELECT id, name, description, price, category_id, created_at FROM products WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, fmt.Errorf("продукт с id %d не найден", id)
		}
		return p, err
	}
	return p, nil
}

func (r *ProductPostgresRepository) Create(p entity.Product) (int, error) {
	var id int
	query := `
		INSERT INTO products (name, description, price, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.DB.QueryRow(query, p.Name, p.Description, p.Price, p.CategoryID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductPostgresRepository) Update(id int, p entity.Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, category_id = $4
		WHERE id = $5
	`
	result, err := r.DB.Exec(query, p.Name, p.Description, p.Price, p.CategoryID, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("продукт с id %d не найден", id)
	}

	return nil
}

func (r *ProductPostgresRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("продукт с id %d не найден", id)
	}

	return nil
}
