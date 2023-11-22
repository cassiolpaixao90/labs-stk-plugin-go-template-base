package out

import (
	"context"
	"{{project_name}}/internal/application/domain"
)

type GetAll{{capitalize_resource}}ClientPort interface {
	GetAll{{capitalize_resource}}(ctx context.Context, id string) (*domain.{{capitalize_resource}}Domain, error)
}
