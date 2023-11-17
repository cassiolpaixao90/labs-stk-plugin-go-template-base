package {{resource}}

import "github.com/labstack/echo/v4"

type Get{{resource|capitalize}}ByIdHandler struct{}

type IGet{{resource|capitalize}}ByIdHandler interface {
	Get{{resource|capitalize}}ById(c echo.Context) error
}

func NewGet{{resource|capitalize}}ByIdIdHandler() IGet{{resource|capitalize}}ByIdHandler {
	return &Get{{resource|capitalize}}ByIdHandler{}
}

func (b *Get{{resource|capitalize}}ByIdHandler) Get{{resource|capitalize}}ById(c echo.Context) error {
	return c.JSON(200, struct{}{})
}
