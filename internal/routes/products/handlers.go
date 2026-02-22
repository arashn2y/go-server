package products

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/json"
	"github.com/arashn0uri/go-server/internal/utils"
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
	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.GetAllProducts)
		r.Post("/", h.CreateProduct)
		r.Get("/{id}", h.GetProductByID)
		r.Put("/{id}", h.UpdateProduct)
		r.Delete("/{id}", h.DeleteProduct)
	})
}

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.service.Products(r.Context())

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	pgID, err := utils.ToPgUUID(productID)

	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := h.service.ProductByID(r.Context(), pgID)

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, product)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req form.CreateProductRequest

	if err := json.Read(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.CreateProduct(r.Context(), req)

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusCreated, map[string]string{"message": "product created successfully"})
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	pgID, err := utils.ToPgUUID(productID)

	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	var req form.UpdateProductRequest

	if err := json.Read(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err = h.service.UpdateProduct(r.Context(), pgID, req)

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, map[string]string{"message": "product updated successfully"})
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	pgID, err := utils.ToPgUUID(productID)

	if err != nil {
		json.WriteError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	rows, err := h.service.DeleteProduct(r.Context(), pgID)

	if rows == 0 {
		json.WriteError(w, http.StatusNotFound, "product not found")
		return
	}

	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.Write(w, http.StatusOK, map[string]string{"message": "product deleted successfully"})
}
