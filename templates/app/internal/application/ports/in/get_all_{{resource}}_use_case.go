package in

import (
	"context"
	"{{project_name}}}}/internal/application/domain"
)

type GetAll{{resource|capitalize}}UseCase interface {
	GetAll{{resource|capitalize}}(ctx context.Context, id string) (*domain.{{resource|capitalize}}Domain, error)
}
