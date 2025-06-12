package entity

import (
	"database/sql"
)

type OrdersItem struct {
	ID        string
	OrderID   string
	ProductID string
	Quantity  sql.NullInt64
	UnitPrice sql.NullFloat64
}
