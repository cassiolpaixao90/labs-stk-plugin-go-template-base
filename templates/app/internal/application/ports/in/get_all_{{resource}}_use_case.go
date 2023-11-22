package in

import (
	"context"
	"{{project_name}}/internal/application/domain"
)

type GetAll{{capitalize_resource}}UseCase interface {
	GetAll{{capitalize_resource}}(ctx context.Context, id string) (*domain.{{capitalize_resource}}Domain, error)
}
