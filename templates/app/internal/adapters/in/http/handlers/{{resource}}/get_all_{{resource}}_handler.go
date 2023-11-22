package {{resource}}

import (
	"context"
	"github.com/labstack/echo/v4"
	"{{project_name}}/internal/application/ports/in"
)

type GetAll{{resource|capitalize}}Handler struct{
	in.GetAll{{resource|capitalize}}UseCase
}

type IGetAll{{resource|capitalize}}Handler interface {
	GetAll{{resource|capitalize}}(c echo.Context) error
}

func NewGetAll{{resource|capitalize}}Handler(getAll{{resource|capitalize}}UseCase in.GetAll{{resource|capitalize}}UseCase) IGetAll{{resource|capitalize}}Handler {
	return &GetAll{{resource|capitalize}}Handler{getAll{{resource|capitalize}}UseCase}
}

func (b *GetAll{{resource|capitalize}}Handler) GetAll{{resource|capitalize}}(c echo.Context) error {
	{{resource}}, err := b.GetAll{{resource|capitalize}}UseCase.GetAll{{resource|capitalize}}(context.Background(), "")
	if err != nil {
		return err
	}
	return c.JSON(200, {{resource}})
}