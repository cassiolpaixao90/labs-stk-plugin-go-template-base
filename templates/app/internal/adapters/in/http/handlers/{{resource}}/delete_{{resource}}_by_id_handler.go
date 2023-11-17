package {{resource}}

import "github.com/labstack/echo/v4"

type Delete{{resource|capitalize}}ByIdHandler struct{}

type IDelete{{resource|capitalize}}ByIdHandler interface {
	Delete{{resource|capitalize}}ById(c echo.Context) error
}

func NewDelete{{resource|capitalize}}ByIdHandler() IDelete{{resource|capitalize}}ByIdHandler {
	return &Delete{{resource|capitalize}}ByIdHandler{}
}

func (b *Delete{{resource|capitalize}}ByIdHandler) Delete{{resource|capitalize}}ById(c echo.Context) error {
	return c.JSON(204, struct{}{})
}
