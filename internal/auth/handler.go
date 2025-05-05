package auth

import (
	"fmt"
	"github.com/jamal23041989/go_short_links/configs"
	"github.com/jamal23041989/go_short_links/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := LoginResponse{
			Token: h.Config.Auth.Secret,
		}
		resp.Json(w, data, 200)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Registration")
	}
}

func NewAuthHandler(r *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	r.HandleFunc("POST /auth/login", handler.Login())
	r.HandleFunc("POST /auth/register", handler.Register())
}
