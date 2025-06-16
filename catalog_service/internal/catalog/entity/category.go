package entity

import "time"

type Category struct {
	ID        int `gorm:"primaryKey;autoIncrement;column:id"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
