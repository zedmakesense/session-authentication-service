package handlers

import "github.com/session-authentication-service/internal/service"

type Deps struct {
	UserService *service.UserService
}

type Handlers struct {
	deps Deps
}

func NewHandlers(d Deps) *Handlers {
	return &Handlers{deps: d}
}
