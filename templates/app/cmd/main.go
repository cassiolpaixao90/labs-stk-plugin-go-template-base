package main

import (
	"{{project_name}}/internal/adapter/in/http"
)

func main() {
	server := http.NewServer()
	server.Start()
}
