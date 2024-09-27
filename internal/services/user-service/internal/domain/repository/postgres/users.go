package postgres

import (
	"context"
	"database/sql"
	"errors"
	"project-management-system/internal/user-service/internal/domain/model"
)

type UserInfoRepo interface {
	Save(ctx context.Context, userId string, userInfo model.UserInfo) error
	WithTx(tx *sql.Tx) *UserInfoRepository
}

type UsersRepository struct {
	db           *sql.DB
	userInfoRepo UserInfoRepo
}

func NewUsersRepository(db *sql.DB, userInfoRepo UserInfoRepo) *UsersRepository {
	return &UsersRepository{db, userInfoRepo}
}

func (r *UsersRepository) findOne(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()

	user := model.User{}
	err = stmt.QueryRowContext(ctx, args...).Scan(
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
		return nil, errors.Join(ErrQueryRowWithContext, err)
	}

	return &user, nil
}

func (r *UsersRepository) FindById(ctx context.Context, id string) (*model.User, error) {
	query := `SELECT 
    	u.id, u.email, u.password, u.created_at, u.updated_at, ui.first_name, ui.last_name, ui.created_at, ui.updated_at
		FROM users AS u LEFT JOIN user_info AS ui ON u.id = ui.user_id WHERE u.id=$1`
	return r.findOne(ctx, query, id)
}

func (r *UsersRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT 
    	u.id, u.email, u.password, u.created_at, u.updated_at, ui.first_name, ui.last_name, ui.created_at, ui.updated_at
		FROM users AS u LEFT JOIN user_info AS ui ON u.id = ui.user_id WHERE u.email=$1`
	return r.findOne(ctx, query, email)
}

func (r *UsersRepository) Save(ctx context.Context, user model.User) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	saveUserQuery := `INSERT INTO users(id, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = tx.ExecContext(ctx, saveUserQuery, user.Id, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return errors.Join(ErrExecWithContext, err)
	}

	err = r.userInfoRepo.WithTx(tx).Save(ctx, user.Id, user.UserInfo)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return ErrCommitTrx
	}

	return nil
}
