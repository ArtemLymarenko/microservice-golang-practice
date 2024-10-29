package postgres

import (
	"database/sql"
	"reflect"
)

// ToNullable retrieves comparable value and converts it to sql null value
// making it invalid if it is zero value otherwise valid
func ToNullable[T comparable](value T) sql.Null[T] {
	v := reflect.ValueOf(value)

	if v.IsZero() {
		return sql.Null[T]{Valid: false}
	}

	return sql.Null[T]{V: value, Valid: true}
}
