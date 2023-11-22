package http

import (
	"github.com/labstack/echo/v4"
	"{{project_name}}/internal/adapters/in/http/handlers"
)

type Routes struct {
	echo *echo.Echo
}

func NewHttpRoutes(e *echo.Echo) *Routes {
	return &Routes{echo: e}
}

func (r *Routes) SetupRouter() {
	r.echo.GET("/{{resource}}", handlers.NewGetAll{{capitalize_resource}}Instance().GetAll{{capitalize_resource}})
}
