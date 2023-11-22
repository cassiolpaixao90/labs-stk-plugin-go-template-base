package services

import (
	"context"
	"{{project_name}}/internal/application/domain"
	"{{project_name}}/internal/application/ports/in"
	"{{project_name}}/internal/application/ports/out"
)

type GetAll{{resource|capitalize}}Service struct {
	GetAll{{resource|capitalize}}ClientPort out.GetAll{{resource|capitalize}}ClientPort
}

func NewGetAll{{resource|capitalize}}Service(
	getAll{{resource|capitalize}}Port out.GetAll{{resource|capitalize}}ClientPort,
) in.GetAll{{resource|capitalize}}UseCase {
	return &GetAll{{resource|capitalize}}Service{getAll{{resource|capitalize}}Port}
}

func (g *GetAll{{resource|capitalize}}Service) GetAll{{resource|capitalize}}(ctx context.Context, id string) (*domain.{{resource|capitalize}}Domain, error) {
	book, err := g.GetAll{{resource|capitalize}}ClientPort.GetAll{{resource|capitalize}}(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
