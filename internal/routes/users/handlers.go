package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"

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
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetAllUsers)
	})
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.Users(r.Context())

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, users)
}
