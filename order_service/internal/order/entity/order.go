package entity

import (
	"time"
)

type Order struct {
	ID            string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID        string
	TotalPrice    float64
	DeliveryPrice float64
	Currency      string
	Status        string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OrderID   string
	ProductID string
	Quantity  int
	UnitPrice float64
}
