package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// Deps lists the concrete resources handlers need.
// Add DB *sql.DB, Cache *redis.Client, Logger, etc.
type Deps struct {
	// DB *sql.DB
	// Cache *redis.Client
	// Logger *zap.Logger
}

type Handlers struct {
	deps Deps
}

func NewHandlers(d Deps) *Handlers { return &Handlers{deps: d} }

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}
