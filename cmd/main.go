package main

import (
	"github.com/XATAB1CH/todo-app"
	"github.com/XATAB1CH/todo-app/pkg/handler"
	"github.com/labstack/gommon/log"
)

func main() {
	handlers := handler.NewHandler()
	server := new(todo.Server)
	err := server.Run(":8080", handlers.InitRoutes())
	if err != nil {
		log.Fatal("Error occurred while starting server: ", err.Error())
	}
}
