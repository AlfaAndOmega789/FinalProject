package handler

import (
	"auth/internal/user/usecase"
	"auth/pkg/jwt"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
	roleID, _ := strconv.Atoi(roleIDStr)

	err = h.UserUC.Register(usecase.RegisterInput{
		Email:    email,
		Password: password,
		Name:     name,
		RoleID:   roleID,
	})
	if err != nil {
		if err == usecase.ErrUserExists {
			http.Error(w, "Пользователь уже существует", http.StatusConflict)
			return
		}
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Пользователь создан")
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

	fmt.Fprintf(w, "User: %s, Email: %s, RoleID: %d", user.Name, user.Email, user.RoleID)
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
