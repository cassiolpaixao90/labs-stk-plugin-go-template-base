package {{resource}}

import (
	"context"
	"github.com/labstack/echo/v4"
	"{{project_name}}/internal/application/ports/in"
)

type GetAll{{custom_resource}}Handler struct{
	in.GetAll{{custom_resource}}UseCase
}

type IGetAll{{custom_resource}}Handler interface {
	GetAll{{custom_resource}}(c echo.Context) error
}

func NewGetAll{{custom_resource}}Handler(getAll{{custom_resource}}UseCase in.GetAll{{custom_resource}}UseCase) IGetAll{{custom_resource}}Handler {
	return &GetAll{{custom_resource}}Handler{getAll{{custom_resource}}UseCase}
}

func (b *GetAll{{custom_resource}}Handler) GetAll{{custom_resource}}(c echo.Context) error {
	{{resource}}, err := b.GetAll{{custom_resource}}UseCase.GetAll{{custom_resource}}(context.Background(), "")
	if err != nil {
		return err
	}
	return c.JSON(200, {{resource}})
}