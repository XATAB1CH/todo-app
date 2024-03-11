package main

import (
	"github.com/XATAB1CH/todo-app"
	"github.com/XATAB1CH/todo-app/pkg/handler"
	"github.com/XATAB1CH/todo-app/pkg/repository"
	"github.com/XATAB1CH/todo-app/pkg/service"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error installing configs: %s", err.Error())
	}

	repos := repository.NewRepository() // зависимости
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	err := server.Run(viper.GetString("8080"), handlers.InitRoutes())
	if err != nil {
		log.Fatal("Error occurred while starting server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
