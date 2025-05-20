package models

import (
	"database/sql"
	"time"
)

type Orders struct {
	ID            int
	UID           string
	TotalPrice    sql.NullFloat64
	DeliveryPrice sql.NullFloat64
	Currency      string
	Status        string
	CreatedAt     time.Time
}
