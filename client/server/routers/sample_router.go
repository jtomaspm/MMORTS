package routers

import (
	"client_server/render"
	"html/template"
	"net/http"
)

func sampleHandler(templates *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title": "Home Page",
			"Body":  "Welcome to the home page!",
		}
		render.RenderTemplate(w, "index", data, templates)
	}
}

func MountSampleRouter(prefix string, templates *template.Template) {
	http.HandleFunc(prefix, sampleHandler(templates))
}
