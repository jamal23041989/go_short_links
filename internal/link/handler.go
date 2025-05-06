package link

import (
	"fmt"
	"github.com/jamal23041989/go_short_links/pkg/req"
	"github.com/jamal23041989/go_short_links/pkg/resp"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := h.LinkRepository.Create(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp.Json(w, createdLink, 201)
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
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

func NewLinkHandler(r *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	r.HandleFunc("POST /link", handler.Create())
	r.HandleFunc("GET /{hash}", handler.GoTo())
	r.HandleFunc("PATCH /link/{id}", handler.Update())
	r.HandleFunc("DELETE /link/{id}", handler.Delete())
}
