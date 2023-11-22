package dtos

import "{{project_name}}/internal/application/domain"

type GetAll{{capitalize_resource}}Response struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (g *GetAll{{capitalize_resource}}Response) ToDomain() *domain.{{capitalize_resource}}Domain {
	return &domain.{{capitalize_resource}}Domain{
		ID:   g.ID,
		Name: g.Name,
	}
}
