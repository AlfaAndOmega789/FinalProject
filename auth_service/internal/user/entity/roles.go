package entity

type Role struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string
}

type Permission struct {
	ID          int    `gorm:"primaryKey"`
	Code        string `gorm:"unique;not null"`
	Description string
}

type RolePermission struct {
	RoleID       int `gorm:"primaryKey"`
	PermissionID int `gorm:"primaryKey"`
}
