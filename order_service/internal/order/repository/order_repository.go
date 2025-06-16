package repository

import (
	"gorm.io/gorm"
	"order/internal/order/entity"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Create(order *entity.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) GetByID(id string) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) Delete(id string) error {
	return r.DB.Delete(&entity.Order{}, "id = ?", id).Error
}

func (r *OrderRepository) Update(id string, status string) error {
	return r.DB.Model(&entity.Order{}).
		Where("id = ?", id).
		Update("status", status).Error
}
