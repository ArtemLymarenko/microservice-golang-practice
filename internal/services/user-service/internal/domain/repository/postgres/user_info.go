package postgres

import (
	"context"
	"database/sql"
	"errors"
	"project-management-system/internal/user-service/internal/domain/model"
)

type DBContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type UserInfoRepository struct {
	db DBContext
}

func NewUserInfoRepository(db *sql.DB) *UserInfoRepository {
	return &UserInfoRepository{db}
}

func (r *UserInfoRepository) Save(ctx context.Context, userId string, userInfo model.UserInfo) error {
	saveUserInfoQuery := `INSERT INTO user_info(user_id, first_name, last_name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, saveUserInfoQuery,
		userId,
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.CreatedAt,
		userInfo.UpdatedAt,
	)

	if err != nil {
		return errors.Join(ErrExecWithContext, err)
	}

	return nil
}

func (r *UserInfoRepository) WithTx(tx *sql.Tx) *UserInfoRepository {
	return &UserInfoRepository{db: tx}
}
