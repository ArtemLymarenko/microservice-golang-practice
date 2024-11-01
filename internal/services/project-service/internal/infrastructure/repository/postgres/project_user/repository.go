package projectUserRepoPostgres

import (
	"database/sql"
	"project-management-system/internal/pkg/sqlStorage"
)

type ProjectUserRepository struct {
	db sqlStorage.DB
}

func New(
	db sqlStorage.DB,
) *ProjectUserRepository {
	return &ProjectUserRepository{
		db: db,
	}
}

func (pu *ProjectUserRepository) WithTx(tx *sql.Tx) *ProjectUserRepository {
	return New(tx)
}
