package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"

	"github.com/arashn0uri/go-server/internal/config"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/routes"
)

type application struct {
	config config.Config
	db     *pgx.Conn
}

func LoadApplication() application {
	cfg := config.Load()
	application := application{
		config: cfg,
	}

	return application
}

func (app *application) Mount() (http.Handler, error) {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is up..."))
	})

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, app.config.DB.DSN)
	if err != nil {
		return nil, err
	} else {
		logrus.Info("connected to the database successfully")
	}

	db := repository.New(conn)
	routes := routes.New(db)

	return routes, nil
}

func (app *application) Run(h http.Handler) error {
	defer app.db.Close(context.Background())
	srv := &http.Server{
		Addr:         app.config.Addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	logrus.Info("Server is running on port", app.config.Addr)

	return srv.ListenAndServe()
}
