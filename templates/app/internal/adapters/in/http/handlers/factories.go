package handlers

import (
	"{{project_name}}/internal/adapters/in/http/handlers/{{resource}}"
	{{resource}}_service "{{project_name}}/internal/adapters/out/http/client/{{resource}}_service"
	"{{project_name}}/internal/application/services"
)

func NewGetAll{{custom_resource}}Instance() {{resource}}.IGetAll{{custom_resource}}Handler {
	client := {{resource}}_service.NewGetAll{{custom_resource}}Client()
	service := services.NewGetAll{{custom_resource}}Service(client)
	return {{resource}}.NewGetAll{{custom_resource}}Handler(service)
}
