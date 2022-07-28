package v1

import (
	users_entity "aero-internship/internal/domain/entity/users"
	"aero-internship/internal/domain/usecase"
	"aero-internship/internal/domain/usecase/auth_usecase"
	"aero-internship/pkg/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type AuthHandler struct {
	service usecase.Service
	cfg     *config.Config
}

func NewAuthHandler(cfg *config.Config, service usecase.Service) *AuthHandler {
	return &AuthHandler{service: service, cfg: cfg}
}

func (a AuthHandler) SignUp(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var userDTO users_entity.UserDTO
	userDTO_json, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Can't read request body: %v", err)))
		return
	}
	err = json.Unmarshal(userDTO_json, &userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Can't unmarshal body: %v", err)))
		return
	}
	tokens, err := a.service.AuthService.RegisterUser(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Can't register user: %v", err)))
		return
	}
	tokensJSON, _ := json.Marshal(tokens)
	refresh_ttl, _ := time.ParseDuration(a.cfg.GetRefreshTTL())
	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		MaxAge:   int(refresh_ttl.Seconds()),
		HttpOnly: true,
	}
	http.SetCookie(w, refreshCookie)
	w.Write(tokensJSON)
}

func (a AuthHandler) SignIn(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	email, password, ok := r.BasicAuth()
	if !ok {

		log.Printf("basic auth went wrong!")

		w.Header().Set("WWW-Authenticate", "Basic")
		w.WriteHeader(http.StatusUnauthorized)

		w.Write([]byte("Auth went wrong!"))
		return
	}
	tokens, err := a.service.AuthService.SignIn(&auth_usecase.SignInDTO{
		Email:    email,
		Password: password,
	})
	if err != nil {
		w.Header().Set("WWW-Authenticate", "Basic")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("Auth went wrong: %v", err)))
		return
	}
	tokensJSON, _ := json.Marshal(tokens)
	refresh_ttl, _ := time.ParseDuration(a.cfg.GetRefreshTTL())
	refreshCookie := &http.Cookie{
		Name:   "refresh_token",
		Value:  tokens.RefreshToken,
		MaxAge: int(refresh_ttl.Seconds()),
	}
	http.SetCookie(w, refreshCookie)
	w.Write(tokensJSON)
}
