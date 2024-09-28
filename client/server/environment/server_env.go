package environment

import (
	"log"
	"os"
)

type ServerEnv struct {
	Port string
}

func LoadServerEnv(default_port string) ServerEnv {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":" + default_port
		log.Println("Using default port...")
	} else {
		port = ":" + port
	}
	return ServerEnv{
		Port: port,
	}
}
