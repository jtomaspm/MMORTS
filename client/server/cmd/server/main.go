package main

import (
	"client_server/environment"
	"client_server/routers"
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Println("Client Server starting...")

	//Load Templates
	templates := template.Must(template.ParseGlob("static/html/*.html"))

	//Routes
	routers.MountStaticRouter("/static/", "./static")
	routers.MountSampleRouter("/", templates)

	server_env := environment.LoadServerEnv("8080")
	err := http.ListenAndServe(server_env.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
