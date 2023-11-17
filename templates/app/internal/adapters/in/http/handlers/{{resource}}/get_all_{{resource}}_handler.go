package {{resource}}

import "github.com/labstack/echo/v4"

type GetAll{{resource|capitalize}}Handler struct{}

type IGetAll{{resource|capitalize}}Handler interface {
	GetAll{{resource|capitalize}}(c echo.Context) error
}

func NewGetAll{{resource|capitalize}}Handler() IGetAll{{resource|capitalize}}Handler {
	return &GetAll{{resource|capitalize}}Handler{}
}

func (b *GetAll{{resource|capitalize}}Handler) GetAll{{resource|capitalize}}(c echo.Context) error {
	return c.JSON(200, struct{}{})
}
