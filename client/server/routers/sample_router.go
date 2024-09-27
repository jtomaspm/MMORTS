package routers

import (
	"client_server/utils"
	"net/http"
)

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home Page",
		"Body":  "Welcome to the home page!",
	}
	utils.RenderTemplate(w, "index", data)
}

func MountSampleRouter(prefix string) {
	http.HandleFunc(prefix, sampleHandler)
}
