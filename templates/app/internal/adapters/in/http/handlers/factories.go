package handlers

import (
	"{{project_name}}/internal/adapters/in/http/handlers/{{resource}}"
	{{resource}}_service "{{project_name}}/internal/adapters/out/http/client/{{resource}}_service"
	"{{project_name}}/internal/application/services"
)

func NewGetAll{{capitalize_resource}}Instance() {{resource}}.IGetAll{{capitalize_resource}}Handler {
	client := {{resource}}_service.NewGetAll{{capitalize_resource}}Client()
	service := services.NewGetAll{{capitalize_resource}}Service(client)
	return {{resource}}.NewGetAll{{capitalize_resource}}Handler(service)
}
