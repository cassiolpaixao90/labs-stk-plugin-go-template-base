package out

import (
	"context"
	"{{project_name}}/internal/application/domain"
)

type GetAll{{custom_resource}}ClientPort interface {
	GetAll{{custom_resource}}(ctx context.Context, id string) (*domain.{{custom_resource}}Domain, error)
}
