// main.go
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vikash-parashar/task-remainder/config"
	"github.com/vikash-parashar/task-remainder/handlers"
)

// TemplateData is a struct to hold data for the template
type TemplateData struct {
	Message string
}

func main() {
	config.InitDB()
	defer config.CloseDB()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// r.Group(func(r chi.Router) {
	// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 		data := TemplateData{
	// 			Message: "Hello from the server!",
	// 		}
	// 		render.RenderTemplate(w, "index", data)
	// 	})
	// 	r.Post("/get-current-user", handlers.GetCurrentUser)
	// 	r.Get("/tasks", handlers.GetTasksForUser)
	// 	r.Get("/tasks/{id}", handlers.GetTaskForUser)
	// 	r.Post("/tasks", handlers.CreateTaskForUser)
	// 	r.Patch("/tasks/{id}", handlers.UpdateTaskForUser)
	// 	r.Delete("/tasks/{id}", handlers.DeleteTaskForUser)
	// })

	// Unprotected routes
	// r.Post("/Register", handlers.RegisterUser)
	// r.Post("/login", handlers.LoginUser)
	r.Get("/page/home", handlers.RenderHomePage)
	r.Get("/page/login", handlers.RenderLoginPage)
	r.Get("/page/register", handlers.RenderRegisterPage)

	http.ListenAndServe(":8080", r)
}
