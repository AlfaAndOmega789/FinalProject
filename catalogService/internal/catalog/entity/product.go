package entity

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description sql.NullString
	Price       float64
	CategoryID  sql.NullInt64
	CreatedAt   time.Time
}

type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
