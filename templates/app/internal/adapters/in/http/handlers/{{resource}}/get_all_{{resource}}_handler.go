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

func NewGetAll{{resource|capitalize}}Handler() IGetAll{{resource|capitalize}}Handler {
	return &GetAll{{resource|capitalize}}Handler{}
}

func (b *GetAll{{resource|capitalize}}Handler) GetAll{{resource|capitalize}}(c echo.Context) error {
	books, err := b.Get{{resource|capitalize}}UseCase.GetAll{{resource|capitalize}}(context.Background(), "")
	if err != nil {
		return err
	}
	return c.JSON(200, {{resource}})
}