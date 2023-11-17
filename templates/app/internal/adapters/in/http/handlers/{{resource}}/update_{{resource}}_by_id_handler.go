package {{resource}}

import "github.com/labstack/echo/v4"

type Update{{resource|capitalize}}ByIdHandler struct{}

type IUpdate{{resource|capitalize}}ByIdHandler interface {
	Update{{resource|capitalize}}ById(c echo.Context) error
}

func NewUpdate{{resource|capitalize}}ByIdHandler() IUpdate{{resource|capitalize}}ByIdHandler {
	return &Update{{resource|capitalize}}ByIdHandler{}
}

func (b *Update{{resource|capitalize}}ByIdHandler) Update{{resource|capitalize}}ById(c echo.Context) error {
	return c.JSON(204, struct{}{})
}
