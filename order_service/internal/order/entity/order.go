package entity

import (
	"time"
)

type Order struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	UserID        int
	TotalPrice    float64
	DeliveryPrice float64
	Currency      string
	Status        string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	UserID        int
	TotalPrice    float64
	DeliveryPrice float64
	Currency      string
	Status        string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}
