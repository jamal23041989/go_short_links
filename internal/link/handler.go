package link

import (
	"net/http"
)

type LinkHandlerDeps struct {
}

type LinkHandler struct {
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func NewLinkHandler(r *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		Config: deps.Config,
	}

	r.HandleFunc("POST /link", handler.Create())
	r.HandleFunc("GET /{alias}", handler.GoTo())
	r.HandleFunc("PATCH /link/{id}", handler.Update())
	r.HandleFunc("DELETE /link/{id}", handler.Delete())
}
