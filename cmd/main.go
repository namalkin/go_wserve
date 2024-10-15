package main

import (
	"log"

	"github.com/namalkin/go_wserve"
	"github.com/namalkin/go_wserve/pkg/handler"
	"github.com/namalkin/go_wserve/pkg/repository"
	"github.com/namalkin/go_wserve/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(go_wserve.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}
