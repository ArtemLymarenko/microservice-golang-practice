package postgres

import (
	"database/sql"
	"project-management-system/internal/user-service/internal/domain/model"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (r *UsersRepository) GetById(id string) *model.User {
	return nil
}

func (r *UsersRepository) Save() error {
	return nil
}
