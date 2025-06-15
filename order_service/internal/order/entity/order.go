package entity

import (
	"time"
)

type Order struct {
	ID            string
	UserID        string
	TotalPrice    float64
	DeliveryPrice float64
	Currency      string
	Status        string
	CreatedAt     time.Time
}

type OrderItem struct {
	ID        string
	OrderID   string
	ProductID string
	Quantity  int
	UnitPrice float64
}
