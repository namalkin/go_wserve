package main

import (
	"log"

	"github.com/namalkin/go_wserve"
	"github.com/namalkin/go_wserve/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)

	srv := new(go_wserve.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}
