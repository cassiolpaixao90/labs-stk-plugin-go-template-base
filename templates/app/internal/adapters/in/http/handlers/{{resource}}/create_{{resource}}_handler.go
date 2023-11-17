package {{resource}}

import "github.com/labstack/echo/v4"

type Create{{resource|capitalize}}Handler struct{}

type ICreate{{resource|capitalize}}Handler interface {
	Create{{resource|capitalize}}(c echo.Context) error
}

func NewCreate{{resource|capitalize}}Handler() ICreate{{resource|capitalize}}Handler {
	return &Create{{resource|capitalize}}Handler{}
}

func (b *Create{{resource|capitalize}}Handler) Create{{resource|capitalize}}(c echo.Context) error {
	return c.JSON(201, struct{}{})
}
