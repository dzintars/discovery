package main

import (
	"log"
	"time"

	"github.com/jakewright/muxinator"
	"github.com/oswee/discovery/controller"
	"github.com/oswee/discovery/domain"
	"github.com/oswee/discovery/service"
)

func main() {
	// port := os.Getenv("SERVICE_PORT")

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

	log.Fatal(router.ListenAndServe(":" + "9000"))
}
