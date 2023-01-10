package main

import (
	"log"
	"todo/handler"
	todo "todo/pkg"
	"todo/repository"
	"todo/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}
