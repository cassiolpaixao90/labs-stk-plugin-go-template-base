package out

import (
	"context"
	"{{project_name}}}}/internal/application/domain"
)

type GetAll{{resource|capitalize}}ClientPort interface {
	GetAll{{resource|capitalize}}(ctx context.Context, id string) (*domain.{{resource|capitalize}}Domain, error)
}
