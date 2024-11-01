package sqlrow

import (
	"database/sql"
	"time"
)

type Project struct {
	Id                string
	Name              string
	Description       sql.Null[string]
	Status            string
	ProductionStartAt sql.Null[time.Time]
	ProductionEndAt   sql.Null[time.Time]
	CreatedAt         sql.Null[time.Time]
	UpdatedAt         sql.Null[time.Time]
	ArchivedAt        sql.Null[time.Time]
}

type UserRole struct {
	UserId string
	Role   string
}
