package main

import (
	"github.com/spf13/viper"
	"iiiproject"
	"iiiproject/pkg/handler"
	"iiiproject/pkg/repository"
	"iiiproject/pkg/service"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error while initializating config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(iiiproject.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
