package services

import (
	"context"
	"{{project_name}}/internal/application/domain"
	"{{project_name}}/internal/application/ports/in"
	"{{project_name}}/internal/application/ports/out"
)

type GetAll{{capitalize_resource}}Service struct {
	GetAll{{capitalize_resource}}ClientPort out.GetAll{{capitalize_resource}}ClientPort
}

func NewGetAll{{capitalize_resource}}Service(
	getAll{{capitalize_resource}}Port out.GetAll{{capitalize_resource}}ClientPort,
) in.GetAll{{capitalize_resource}}UseCase {
	return &GetAll{{capitalize_resource}}Service{getAll{{capitalize_resource}}Port}
}

func (g *GetAll{{capitalize_resource}}Service) GetAll{{capitalize_resource}}(ctx context.Context, id string) (*domain.{{capitalize_resource}}Domain, error) {
	book, err := g.GetAll{{capitalize_resource}}ClientPort.GetAll{{capitalize_resource}}(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
