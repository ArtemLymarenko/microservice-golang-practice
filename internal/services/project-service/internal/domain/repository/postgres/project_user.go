package postgres

import (
	"context"
	"database/sql"
)

type projectUserDB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type ProjectUserRepository struct {
	db projectUserDB
}

func NewProjectUserRepository(db projectUserDB) *ProjectUserRepository {
	return &ProjectUserRepository{db}
}

func (p *ProjectUserRepository) Save(ctx context.Context, projectId string, userId string) error {
	saveProjectUserQuery := `INSERT INTO 
    	projects_users("project_id", "user_id")
		VALUES ($1, $2)`

	_, err := p.db.ExecContext(ctx, saveProjectUserQuery, projectId, userId)
	return err
}

func (p *ProjectUserRepository) WithTx(tx *sql.Tx) *ProjectUserRepository {
	return &ProjectUserRepository{tx}
}
