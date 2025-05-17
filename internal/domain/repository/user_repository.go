package repository

import (
	"context"

	"github.com/DevAntonioJorge/go-notes/internal/domain/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) SaveUser(user *models.User) error {
	query := `INSERT INTO users (id, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`
	if _, err := r.db.Exec(context.Background(), query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = $1`
	user := new(models.User)
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err == pgx.ErrNoRows || err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE id = $1`
	user := new(models.User)
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err == pgx.ErrNoRows || err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByName(name string) (*models.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE name = $1`
	user := new(models.User)
	err := r.db.QueryRow(context.Background(), query, name).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err == pgx.ErrNoRows || err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdatePassword(user *models.User, password string) error {
	query := `UPDATE users SET password = $1 WHERE id = $2`
	_, err := r.db.Exec(context.Background(), query, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}
