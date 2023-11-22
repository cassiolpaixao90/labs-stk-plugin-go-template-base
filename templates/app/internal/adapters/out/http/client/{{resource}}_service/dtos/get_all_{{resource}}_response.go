package dtos

import "{{project_name}}/internal/application/domain"

type GetAll{{resource|capitalize}}Response struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func (g *GetAll{{resource|capitalize}}Response) ToDomain() *domain.{{resource|capitalize}}Domain {
	return &domain.{{resource|capitalize}}Domain{
		ID:   g.ID,
		Name: g.Name,
	}
}
