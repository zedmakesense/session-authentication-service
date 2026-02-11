package service

import (
	"context"
	"errors"

	"github.com/session-authentication-service/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, name string, username string, password string) (int64, error) {
	if len(password) < 8 {
		return 0, errors.New("password must be at least 8 characters")
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateUser(ctx, name, username, hashedPassword)
}

func (s *UserService) Login(ctx context.Context, username string, password string) error {
	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	return comparePassword(user.HashedPassword, password)
}
