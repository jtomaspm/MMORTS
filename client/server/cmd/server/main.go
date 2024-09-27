package main

import (
	"client_server/routers"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello World!")

	//Static Dir
	staticDir := "/static/"
	http.Handle(staticDir, http.StripPrefix(staticDir, http.FileServer(http.Dir("static"))))

	//Routes
	routers.MountSampleRouter("/")

	port := ":80"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
