package handlers

import (
	"{{project_name}}/internal/adapters/in/http/handlers/{{resource}}"
	{{resource}}_service "{{project_name}}/internal/adapters/out/http/client/{{resource}}_service"
	"{{project_name}}/internal/application/services"
)

func NewGetAll{{resource|capitalize}}Instance() {{resource}}.IGetAll{{resource|capitalize}}Handler {
	client := {{resource}}_service.NewGetAll{{resource|capitalize}}Client()
	service := services.NewGetAll{{resource|capitalize}}Service(client)
	return {{resource}}.NewGetAll{{resource|capitalize}}Handler(service)
}
