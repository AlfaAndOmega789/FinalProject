package entity

import "time"

type Product struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	CategoryID  int       `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
