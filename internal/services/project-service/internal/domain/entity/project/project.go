package project

import "time"

type Status string

const (
	Idle   Status = "idle"
	Active Status = "active"
	Closed Status = "closed"
)

type Id string

type Name string

type Description string

type Project struct {
	Id                Id
	Name              Name
	Description       Description
	Status            Status
	ProductionStartAt time.Time
	ProductionEndAt   time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArchivedAt        time.Time
}
