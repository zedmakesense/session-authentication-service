package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID             int64
	Name           string
	Username       string
	HashedPassword string
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, name string, username string, hashedPassword string) (int64, error) {
	query := "INSERT INTO users (name, username, hashed_password) VALUES ($1, $2, $3) RETURNING id"

	var id int64

	err := r.db.QueryRow(ctx, query, name, username, hashedPassword).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := "SELECT id, name, username, hashed_password FROM users WHERE username = $1"

	var user User

	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Name, &user.Username, &user.HashedPassword)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
