package main

import (
	"github.com/XATAB1CH/todo-app"
	"github.com/XATAB1CH/todo-app/pkg/handler"
	"github.com/XATAB1CH/todo-app/pkg/repository"
	"github.com/XATAB1CH/todo-app/pkg/service"
	"github.com/labstack/gommon/log"
)

func main() {

	repos := repository.NewRepository() // зависимости
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	err := server.Run(":8080", handlers.InitRoutes())
	if err != nil {
		log.Fatal("Error occurred while starting server: ", err.Error())
	}
}
