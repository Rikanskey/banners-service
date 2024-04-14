package http

import (
	"banners-service/internal/app"
	v1 "banners-service/internal/port/http/v1"
	"banners-service/pkg/logging"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
)

func NewHandler(app app.Application) http.Handler {
	apiRouter := chi.NewRouter()
	addMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/v1", v1.NewHandler(app, apiRouter))

	return rootRouter
}

func addMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logging.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)
	addCORSMiddleware(router)
	//router.Use(mock.MockAuthHTTPMiddleware)
	router.Use(middleware.NoCache)
}

const maxAge = 300

func addCORSMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           maxAge,
	})
	router.Use(corsMiddleware.Handler)
}
