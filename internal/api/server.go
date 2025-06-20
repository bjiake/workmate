package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	"time"
	usrHdl "workmate/internal/api/handler/task"
	"workmate/internal/config"
)

type ServerHTTP struct {
	router http.Handler
}

func NewServerHTTP(
	cfg *config.Config,
	userHandler *usrHdl.TaskHandler,
) *ServerHTTP {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		userHandler.NewTaskHandler(r)

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://%s:%s/api/v1/swagger/doc.json", cfg.Server.Host, cfg.Server.Port)),
		))
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"}, // Замените на ваш источник
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
	handler := c.Handler(r)

	return &ServerHTTP{router: handler}
}

func (sh *ServerHTTP) Start(cfg *config.Config, log *slog.Logger) {
	fmt.Print(fmt.Sprintf("Port is %s", cfg.Server.Port))
	log.Info(fmt.Sprintf("Starting server on port: %s", cfg.Server.Port))
	addr := cfg.Server.Addr + ":" + cfg.Server.Port
	err := http.ListenAndServe(addr, sh.router)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
