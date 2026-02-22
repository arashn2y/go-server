package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/routes/auth"
	"github.com/arashn0uri/go-server/internal/routes/products"
	"github.com/arashn0uri/go-server/internal/routes/users"
)

func New(db *repository.Queries) http.Handler {
	r := chi.NewRouter()

	// auth
	authService := auth.NewService(db)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(r)

	// user
	userService := users.NewService(db)
	userHandler := users.NewHandler(userService)
	userHandler.RegisterRoutes(r)

	// product
	productService := products.NewService(db)
	productHandler := products.NewHandler(productService)
	productHandler.RegisterRoutes(r)

	return r
}
