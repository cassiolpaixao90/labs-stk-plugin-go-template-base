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
	r.echo.GET("/{{resource}}", handlers.NewGetAll{{resource|capitalize}}Instance().GetAll{{resource|capitalize}})
	r.echo.POST("/{{resource}}", handlers.NewCreate{{resource|capitalize}}Instance().Create{{resource|capitalize}})
	r.echo.GET("/{{resource}}/:id", handlers.NewGet{{resource|capitalize}}ByIdInstance().Get{{resource|capitalize}}ById)
	r.echo.PUT("/{{resource}}/:id", handlers.NewUpdate{{resource|capitalize}}ByIdInstance().Update{{resource|capitalize}}ById)
	r.echo.DELETE("/{{resource}}/:id", handlers.NewDelete{{resource|capitalize}}ByIdInstance().Delete{{resource|capitalize}}ById)
}
