package project

import (
	_ "github.com/go-playground/validator/v10"
	"time"
)

type Status string

const (
	StatusIdle   Status = "idle"
	StatusActive Status = "active"
	StatusClosed Status = "closed"
)

type Id string

type Name string

type Description string

type Project struct {
	Id                Id          `validate:"required,uuid"`
	Name              Name        `validate:"required,min=3,max=20"`
	Description       Description `validate:"min=20,max=300"`
	Status            Status      `validate:"required"`
	ProductionStartAt time.Time   `validate:""`
	ProductionEndAt   time.Time   `validate:""`
	CreatedAt         time.Time   `validate:"required"`
	UpdatedAt         time.Time   `validate:"required"`
	ArchivedAt        time.Time   `validate:""`
}
