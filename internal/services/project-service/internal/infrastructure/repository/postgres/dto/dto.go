package postgresdto

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
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArchivedAt        sql.Null[time.Time]
}

type MemberDTO struct {
}
