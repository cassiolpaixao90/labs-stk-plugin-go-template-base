package {{resource}}

import (
	"context"
	"github.com/labstack/echo/v4"
	"{{project_name}}/internal/application/ports/in"
)

type GetAll{{capitalize_resource}}Handler struct{
	in.GetAll{{capitalize_resource}}UseCase
}

type IGetAll{{capitalize_resource}}Handler interface {
	GetAll{{capitalize_resource}}(c echo.Context) error
}

func NewGetAll{{capitalize_resource}}Handler(getAll{{capitalize_resource}}UseCase in.GetAll{{capitalize_resource}}UseCase) IGetAll{{capitalize_resource}}Handler {
	return &GetAll{{capitalize_resource}}Handler{getAll{{capitalize_resource}}UseCase}
}

func (b *GetAll{{capitalize_resource}}Handler) GetAll{{capitalize_resource}}(c echo.Context) error {
	{{resource}}, err := b.GetAll{{capitalize_resource}}UseCase.GetAll{{capitalize_resource}}(context.Background(), "")
	if err != nil {
		return err
	}
	return c.JSON(200, {{resource}})
}