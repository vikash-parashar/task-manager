// main.go
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yourusername/yourproject/config"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	r := chi.NewRouter()

	r.Get("/tasks", GetTasks)
	r.Get("/tasks/{id}", GetTask)
	r.Post("/tasks", CreateTask)
	r.Put("/tasks/{id}", UpdateTask)
	r.Delete("/tasks/{id}", DeleteTask)

	http.ListenAndServe(":8080", r)
}
