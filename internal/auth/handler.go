package auth

import (
	"github.com/jamal23041989/go_short_links/configs"
	"github.com/jamal23041989/go_short_links/pkg/req"
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
		_, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}

		data := LoginResponse{
			Token: h.Config.Auth.Secret,
		}
		resp.Json(w, data, 200)
	}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		data := RegisterResponse{
			Token: h.Config.Auth.Secret,
		}
		resp.Json(w, data, 200)
	}
}

func NewAuthHandler(r *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	r.HandleFunc("POST /auth/login", handler.Login())
	r.HandleFunc("POST /auth/register", handler.Register())
}
