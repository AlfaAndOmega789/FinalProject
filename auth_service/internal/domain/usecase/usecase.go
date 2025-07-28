package usecase

import (
	entity2 "auth/internal/domain/entity"
	"auth/internal/domain/repository"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string
	Password string
	Name     string
	RoleID   uuid.UUID
}

var (
	ErrUserExists        = errors.New("user already exists")
	ErrInvalidCredential = errors.New("invalid credentials")
	ErrRoleNotFound      = errors.New("role not found")
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(r repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (uc *UserUsecase) Register(input RegisterInput) error {
	existing, _ := uc.repo.GetByEmail(input.Email)
	if existing != nil {
		return ErrUserExists
	}

	role, err := uc.repo.GetRoleByID(input.RoleID)
	if err != nil {
		return err
	}
	if role == nil {
		return ErrRoleNotFound
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &entity2.User{
		ID:           uuid.New(),
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
		Name:         input.Name,
		RoleID:       input.RoleID,
		CreatedAt:    time.Now(),
	}

	return uc.repo.Create(newUser)
}

func (uc *UserUsecase) Login(email, password string) (*entity2.User, error) {
	user, err := uc.repo.GetByEmail(email)
	if err != nil || user == nil {
		return nil, ErrInvalidCredential
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredential
	}

	return user, nil
}

func (uc *UserUsecase) GetByID(id string) (*entity2.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUsecase) GetRoleByID(id uuid.UUID) (*entity2.Role, error) {
	return uc.repo.GetRoleByID(id)
}
