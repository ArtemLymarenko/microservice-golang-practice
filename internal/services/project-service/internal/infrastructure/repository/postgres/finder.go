package postgres

import (
	"context"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

func FindOne[T any](
	ctx context.Context,
	db DB,
	scan func(row Scanner) (T, error),
	query string,
	args ...interface{},
) (*T, error) {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	row := stmt.QueryRowContext(ctx, args...)
	result, err := scan(row)
	if err != nil {
		return nil, ErrRowsNotRead
	}

	return &result, nil
}

func FindMany[T any](
	ctx context.Context,
	db DB,
	scan func(rows Scanner) (T, error),
	query string,
	args ...interface{},
) (data []T, err error) {
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return data, ErrRowsNotRead
	}

	for rows.Next() {
		result, err := scan(rows)
		if err != nil {
			return nil, ErrRowsNotRead
		}

		data = append(data, result)
	}

	return data, nil
}
