package main

import (
	"log"

	server "github.com/balatsanandrey25/PWG"
	"github.com/balatsanandrey25/PWG/pkg/handler"
)

func main() {
	handlers := handler.NewHandler()
	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitHandler()); err != nil {
		log.Fatal("The server fell to the death of the brave: %s", err)
	}
}
