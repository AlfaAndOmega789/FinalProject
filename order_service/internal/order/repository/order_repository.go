package repository

import (
	"database/sql"
	"fmt"
	"order/internal/order/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *entity.Order) error {
	query := `
		INSERT INTO orders (id, user_id, total_price, delivery_price, currency, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, now())
	`
	_, err := r.DB.Exec(query,
		order.ID,
		order.UserID,
		order.TotalPrice,
		order.DeliveryPrice,
		order.Currency,
		order.Status,
	)
	return err
}

func (r *OrderRepository) GetByID(id string) (*entity.Order, error) {
	query := `
		SELECT id, user_id, total_price, delivery_price, currency, status, created_at
		FROM orders
		WHERE id = $1
	`

	row := r.DB.QueryRow(query, id)
	order := &entity.Order{}

	err := row.Scan(
		&order.ID,
		&order.UserID,
		&order.TotalPrice,
		&order.DeliveryPrice,
		&order.Currency,
		&order.Status,
		&order.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("заказ не найден")
	}
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Delete(id string) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *OrderRepository) Update(id string, status string) error {
	query := `UPDATE orders SET status = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, status, id)
	return err
}
