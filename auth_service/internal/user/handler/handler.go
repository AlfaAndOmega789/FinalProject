package handler

import (
	"auth/internal/user/usecase"
	"auth/pkg/jwt"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
)

type AuthHandler struct {
	UserUC *usecase.UserUsecase
}

func NewAuthHandler(uc *usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{UserUC: uc}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	formData := parseForm(string(body))
	email := formData["email"]
	password := formData["password"]
	name := formData["name"]
	roleIDStr := formData["role_id"]

	roleUUID, err := uuid.Parse(roleIDStr)
	if err != nil {
		http.Error(w, "Некорректный UUID роли", http.StatusBadRequest)
		return
	}

	err = h.UserUC.Register(usecase.RegisterInput{
		Email:    email,
		Password: password,
		Name:     name,
		RoleID:   roleUUID,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	formData := parseForm(string(body))
	email := formData["email"]
	password := formData["password"]

	user, err := h.UserUC.Login(email, password)
	if err != nil {
		http.Error(w, "Неверные данные", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := jwt.GenerateTokens(user.ID.String())
	if err != nil {
		http.Error(w, "Ошибка токена", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "AccessToken: %s\nRefreshToken: %s", accessToken, refreshToken)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Token required", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := jwt.ParseToken(token, true)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	user, err := h.UserUC.GetByID(userID)
	if err != nil || user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	role, _ := h.UserUC.GetRoleByID(user.RoleID)

	roleName := "неизвестна"
	if role != nil {
		roleName = role.Name
	}

	fmt.Fprintf(w, "User: %s, Email: %s, Role: %s", user.Name, user.Email, roleName)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	formData := parseForm(string(body))
	refreshToken := formData["refresh_token"]

	userID, err := jwt.ParseToken(refreshToken, false)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	accessToken, newRefreshToken, err := jwt.GenerateTokens(userID)
	if err != nil {
		http.Error(w, "Ошибка токена", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "AccessToken: %s\nRefreshToken: %s", accessToken, newRefreshToken)
}

func parseForm(body string) map[string]string {
	form := make(map[string]string)
	pairs := strings.Split(body, "&")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			form[kv[0]] = kv[1]
		}
	}
	return form
}
