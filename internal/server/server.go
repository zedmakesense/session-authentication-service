package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/session-authentication-service/internal/handlers"
)

type Server struct {
	port     int
	handlers *handlers.Handlers
}

func NewServer() *http.Server {
	port := 8080
	if p := os.Getenv("PORT"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			port = v
		}
	}

	deps := handlers.Deps{
		// DB: db,
		// Cache: cache,
		// Logger: logger,
	}
	h := handlers.NewHandlers(deps)

	s := &Server{
		port:     port,
		handlers: h,
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return srv
}
