package repository

import (
	entity2 "auth/internal/domain/entity"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity2.User) error
	GetByEmail(email string) (*entity2.User, error)
	GetByID(id string) (*entity2.User, error)
	GetRoleByID(id uuid.UUID) (*entity2.Role, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *entity2.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) GetByEmail(email string) (*entity2.User, error) {
	var user entity2.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *userRepo) GetByID(id string) (*entity2.User, error) {
	var user entity2.User
	err := r.db.First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *userRepo) GetRoleByID(id uuid.UUID) (*entity2.Role, error) {
	var role entity2.Role
	err := r.db.First(&role, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, err
}
