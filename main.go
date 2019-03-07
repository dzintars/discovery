package main

import (
	"log"
	"os"
	"time"

	"github.com/jakewright/muxinator"
	"github.com/oswee/srp/discovery/controller"
	"github.com/oswee/srp/discovery/domain"
	"github.com/oswee/srp/discovery/service"
)

func main() {
	port := os.Getenv("SERVICE_PORT")

	config := domain.Config{}

	configService := service.ConfigService{
		Config:   &config,
		Location: "config.yaml",
	}

	// Set interwal to fetch config changes from config.yaml
	go configService.Watch(time.Second * 30)

	c := controller.Controller{
		Config: &config,
	}

	router := muxinator.NewRouter()
	router.Get("/read/{serviceName}", c.ReadConfig)

	log.Fatal(router.ListenAndServe(":" + port))
}
