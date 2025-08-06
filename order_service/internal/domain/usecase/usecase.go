package usecase

import "order/internal/domain/entity"

type OrderRepository interface {
	GetByID(id string) (*entity.Order, error)
	Create(order *entity.Order) error
	UpdateStatus(id string, status string) error
	Delete(id string) error
}

type OrderUsecase struct {
	repo OrderRepository
}

func NewOrderUsecase(r OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: r}
}

func (u *OrderUsecase) Create(order *entity.Order) error {
	return u.repo.Create(order)
}

func (u *OrderUsecase) GetByID(id string) (*entity.Order, error) {
	return u.repo.GetByID(id)
}

func (u *OrderUsecase) Delete(id string) error {
	return u.repo.Delete(id)
}

func (u *OrderUsecase) UpdateStatus(id, status string) error {
	return u.repo.UpdateStatus(id, status)
}
