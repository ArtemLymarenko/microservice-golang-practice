package project

import (
	"errors"
	"time"
)

var (
	ErrInvalidStatusValue         = errors.New("provided status value is invalid")
	ErrInvalidProductionStartDate = errors.New("invalid start project date, should be before end date")
	ErrInvalidProductionEndDate   = errors.New("invalid end project date, should be after start date")
	ErrInvalidArchivedAtDate      = errors.New("invalid archived at date")
)

type Validator interface {
	Struct(s interface{}) error
}

func (p *Project) ValidateStatus() error {
	switch p.Status {
	case StatusIdle, StatusActive, StatusClosed:
		return nil
	default:
		return ErrInvalidStatusValue
	}
}

func (p *Project) ValidateProductionStartAt() error {
	if p.ProductionEndAt.IsZero() {
		return nil
	}

	now := time.Now()
	if p.ProductionStartAt.Before(now) || p.ProductionStartAt.After(p.ProductionEndAt) {
		return ErrInvalidProductionStartDate
	}

	return nil
}

func (p *Project) ValidateProductionEndAt() error {
	if p.ProductionEndAt.IsZero() {
		return nil
	}

	now := time.Now()
	if p.ProductionStartAt.Before(now) || p.ProductionEndAt.Before(p.ProductionStartAt) {
		return ErrInvalidProductionEndDate
	}

	return nil
}

func (p *Project) ValidateArchivedAt() error {
	if p.ArchivedAt.IsZero() {
		return nil
	}

	now := time.Now()
	if p.ArchivedAt.Before(now) {
		return ErrInvalidArchivedAtDate
	}

	return nil
}

func (p *Project) Validate(validator Validator) error {
	if err := p.ValidateStatus(); err != nil {
		return err
	}

	if err := p.ValidateProductionStartAt(); err != nil {
		return err
	}

	if err := p.ValidateProductionEndAt(); err != nil {
		return err
	}

	if err := p.ValidateArchivedAt(); err != nil {
		return err
	}

	return validator.Struct(p)
}
