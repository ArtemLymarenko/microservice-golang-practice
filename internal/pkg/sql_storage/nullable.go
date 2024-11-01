package sqlStorage

import (
	"database/sql"
	"reflect"
)

// ComparableToSqlNull retrieves comparable value and converts it to sql null value
// making it invalid if it is zero value otherwise valid
func ComparableToSqlNull[T comparable](value T) sql.Null[T] {
	v := reflect.ValueOf(value)

	if v.IsZero() {
		return sql.Null[T]{Valid: false}
	}

	return sql.Null[T]{V: value, Valid: true}
}
