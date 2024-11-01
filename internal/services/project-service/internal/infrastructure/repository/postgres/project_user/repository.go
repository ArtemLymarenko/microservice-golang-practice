package projectUserRepoPostgres

import (
	"database/sql"
	sqlSstorage "project-management-system/internal/pkg/sql_storage"
)

type ProjectUserRepository struct {
	db sqlSstorage.DB
}

func New(
	db sqlSstorage.DB,
) *ProjectUserRepository {
	return &ProjectUserRepository{
		db: db,
	}
}

func (pu *ProjectUserRepository) WithTx(tx *sql.Tx) *ProjectUserRepository {
	return New(tx)
}
