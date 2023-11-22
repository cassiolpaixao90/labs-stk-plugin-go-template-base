package services

import (
	"context"
	"{{project_name}}/internal/application/domain"
	"{{project_name}}/internal/application/ports/in"
	"{{project_name}}/internal/application/ports/out"
)

type GetAll{{custom_resource}}Service struct {
	GetAll{{custom_resource}}ClientPort out.GetAll{{custom_resource}}ClientPort
}

func NewGetAll{{custom_resource}}Service(
	getAll{{custom_resource}}Port out.GetAll{{custom_resource}}ClientPort,
) in.GetAll{{custom_resource}}UseCase {
	return &GetAll{{custom_resource}}Service{getAll{{custom_resource}}Port}
}

func (g *GetAll{{custom_resource}}Service) GetAll{{custom_resource}}(ctx context.Context, id string) (*domain.{{custom_resource}}Domain, error) {
	book, err := g.GetAll{{custom_resource}}ClientPort.GetAll{{custom_resource}}(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
