package render

import (
	"html/template"
	"net/http"
)

// renderTemplate renders an HTML template with the provided data
func RenderTemplate(w http.ResponseWriter, t string, data interface{}) {
	// Parse and execute the HTML template file
	tmplFile := "templates/" + t + ".page.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the provided data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
