package dtos

import "{{project_name}}/internal/application/domain"

type GetAll{{custom_resource}}Response struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (g *GetAll{{custom_resource}}Response) ToDomain() *domain.{{custom_resource}}Domain {
	return &domain.{{custom_resource}}Domain{
		ID:   g.ID,
		Name: g.Name,
	}
}
