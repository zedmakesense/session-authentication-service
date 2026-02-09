package handlers

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

func NewHandlers(d Deps) *Handlers {
	return &Handlers{deps: d}
}
