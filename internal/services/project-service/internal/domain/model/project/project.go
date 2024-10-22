package project

import "time"

type Status string

const (
	Idle   Status = "idle"
	Active Status = "active"
	Closed Status = "closed"
)

type Project struct {
	Id                string
	Name              string
	Description       string
	Status            Status
	ProductionStartAt time.Time
	ProductionEndAt   time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArchivedAt        time.Time
}
