package handlers

import (
	"{{project_name}}/internal/adapters/in/http/handlers/{{resource}}"
)

func NewGetAll{{resource|capitalize}}Instance() {{resource}}.IGetAll{{resource|capitalize}}Handler {
	client := {{resource}}_service.NewGetAll{{resource|capitalize}}Client()
	service := services.NewGetAll{{resource|capitalize}}Service(client)
	return {{resource}}.NewGetAll{{resource|capitalize}}Handler(service)
}

func NewCreate{{resource|capitalize}}Instance() {{resource}}.ICreate{{resource|capitalize}}Handler {
	return {{resource}}.NewCreate{{resource|capitalize}}Handler()
}

func NewGet{{resource|capitalize}}ByIdInstance() {{resource}}.IGet{{resource|capitalize}}ByIdHandler {
	return {{resource}}.NewGet{{resource|capitalize}}ByIdIdHandler()
}

func NewUpdate{{resource|capitalize}}ByIdInstance() {{resource}}.IUpdate{{resource|capitalize}}ByIdHandler {
	return {{resource}}.NewUpdate{{resource|capitalize}}ByIdHandler()
}

func NewDelete{{resource|capitalize}}ByIdInstance() {{resource}}.IDelete{{resource|capitalize}}ByIdHandler {
	return {{resource}}.NewDelete{{resource|capitalize}}ByIdHandler()
}
