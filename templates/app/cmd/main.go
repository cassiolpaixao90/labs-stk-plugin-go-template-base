package main

import (
	"{{project_name}}/internal/adapters/in/http"
)

func main() {
	server := http.NewServer()
	server.Start()
}
