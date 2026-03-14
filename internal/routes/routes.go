package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arashn0uri/go-server/internal/middleware"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/routes/auth"
	"github.com/arashn0uri/go-server/internal/routes/permissions"
	"github.com/arashn0uri/go-server/internal/routes/products"
	"github.com/arashn0uri/go-server/internal/routes/users"
)

func New(db *repository.Queries) http.Handler {
	r := chi.NewRouter()

	// public routes
	// auth
	authService := auth.NewService(db)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(r)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.Auth)
		// permission
		permissionService := permissions.NewService(db)
		// user
		userService := users.NewService(db, permissionService)
		userHandler := users.NewHandler(userService)
		userHandler.RegisterRoutes(r)

		// product
		productService := products.NewService(db, permissionService)
		productHandler := products.NewHandler(productService)
		productHandler.RegisterRoutes(r)
	})

	return r
}
