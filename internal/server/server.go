package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/session-authentication-service/internal/db"
	"github.com/session-authentication-service/internal/handlers"
	"github.com/session-authentication-service/internal/repository"
	"github.com/session-authentication-service/internal/service"
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

	dsn := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := db.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	deps := handlers.Deps{
		UserService: userService,
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
