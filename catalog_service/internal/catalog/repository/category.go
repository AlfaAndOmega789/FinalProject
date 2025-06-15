package repository

import (
	"catalog/internal/catalog/entity"
	"database/sql"
	"fmt"
)

type CategoryPostgresRepository struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryPostgresRepository {
	return &CategoryPostgresRepository{DB: db}
}

func (r *CategoryPostgresRepository) GetAll() ([]entity.Category, error) {
	rows, err := r.DB.Query("SELECT id, name, created_at FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category

	for rows.Next() {
		var p entity.Category
		err := rows.Scan(&p.ID, &p.Name, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, p)
	}

	return categories, nil
}
func (r *CategoryPostgresRepository) Create(p entity.Category) (int, error) {
	var id int
	query := `
		INSERT INTO categories (name) VALUES ($1) RETURNING id
	`
	err := r.DB.QueryRow(query, p.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CategoryPostgresRepository) Update(id int, p entity.Category) error {
	query := `
		UPDATE categories SET name = $1 WHERE id = $2
	`
	result, err := r.DB.Exec(query, p.Name, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("категория с id %d не найдена", id)
	}

	return nil
}

func (r *CategoryPostgresRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("категория с id %d не найдена", id)
	}

	return nil
}
