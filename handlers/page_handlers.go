package handlers

import (
	"net/http"

	"github.com/vikash-parashar/task-remainder/render"
)

func RenderHomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index", nil)
}
func RenderLoginPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login", nil)
}
func RenderRegisterPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register", nil)
}
