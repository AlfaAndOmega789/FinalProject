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
	if r.Method != http.MethodPost {
		http.Error(w, "не разрешено", http.StatusMethodNotAllowed)
		return
	}

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
			http.Error(w, "пользователь уже существует", http.StatusConflict)
			return
		}
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Пользователь успешно создан")
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Не разрешено", http.StatusMethodNotAllowed)
		return
	}

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
		http.Error(w, "Ошибка в данных", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := jwt.GenerateTokens(user.ID.String())
	if err != nil {
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "AccessToken: %s\nRefreshToken: %s", accessToken, refreshToken)
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
