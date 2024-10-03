package model

import "time"

type ProjectStatus string

const (
	ProjectIdle   ProjectStatus = "idle"
	ProjectActive ProjectStatus = "active"
	ProjectClosed ProjectStatus = "closed"
)

type Project struct {
	Id                string
	Name              string
	Description       string
	Status            ProjectStatus
	ProductionStartAt time.Time
	ProductionEndAt   time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ArchivedAt        time.Time
}
