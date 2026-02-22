package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.Register)
	})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req form.Register

	if err := json.Read(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	users, err := h.service.Register(r.Context(), req)

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, users)
}
