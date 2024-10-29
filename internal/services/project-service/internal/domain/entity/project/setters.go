package project

import "time"

func (p *Project) SetId(id Id) {
	p.Id = id
}

func (p *Project) SetName(name Name) {
	p.Name = name
}

func (p *Project) SetDescription(description Description) {
	p.Description = description
}

func (p *Project) SetStatus(status Status) {
	p.Status = status
}

func (p *Project) SetProductionStartAt(productionStartAt time.Time) {
	p.ProductionStartAt = productionStartAt
}

func (p *Project) SetProductionEndAt(productionEndAt time.Time) {
	p.ProductionEndAt = productionEndAt
}

func (p *Project) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt
}

func (p *Project) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt
}

func (p *Project) SetArchivedAt(archivedAt time.Time) {
	p.ArchivedAt = archivedAt
}
