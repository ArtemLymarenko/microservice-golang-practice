package dto

import "time"

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProjectRequest struct {
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	Status            string    `json:"status"`
	ProductionStartAt time.Time `json:"productionStartAt"`
	ProductionEndAt   time.Time `json:"productionEndAt"`
}
