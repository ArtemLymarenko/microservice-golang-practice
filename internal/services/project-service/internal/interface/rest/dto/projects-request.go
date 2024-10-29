package dto

import "time"

type CreateProjectRequest struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	ProductionStartAt time.Time `json:"productionStartAt,omitempty"`
	ProductionEndAt   time.Time `json:"productionEndAt,omitempty"`
}

type UpdateProjectRequest struct {
	Id                string    `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	ProductionStartAt time.Time `json:"productionStartAt"`
	ProductionEndAt   time.Time `json:"productionEndAt"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	ArchivedAt        time.Time `json:"archivedAt"`
}
