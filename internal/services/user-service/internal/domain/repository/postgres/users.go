package postgres

import (
	"context"
	"database/sql"
	"project-management-system/internal/user-service/internal/domain/model"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (r *UsersRepository) findOne(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := model.User{}
	err = stmt.QueryRowContext(ctx, args).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.UserInfo.FirstName,
		&user.UserInfo.LastName,
		&user.UserInfo.CreatedAt,
		&user.UserInfo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT 
    	u.id, u.email, u.password, u.created_at, u.updated_at, ui.first_name, ui.last_name, ui.created_at, ui.updated_at
		FROM users AS u LEFT JOIN user_info AS ui ON u.id = ui.user_id WHERE u.id=$1`
	return r.findOne(ctx, query, id)
}

func (r *UsersRepository) Save(ctx context.Context, user *model.User) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	saveUserQuery := `INSERT INTO users(id, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, saveUserQuery, user.Id, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	saveUserInfoQuery := `INSERT INTO users(user_id, first_name, last_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, saveUserInfoQuery,
		user.Id,
		user.UserInfo.FirstName,
		user.UserInfo.LastName,
		user.UserInfo.CreatedAt,
		user.UserInfo.UpdatedAt,
	)

	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
