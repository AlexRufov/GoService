package main

import (
	"GoService/app/internal/app"
	"GoService/app/internal/config"
	"GoService/app/pkg/logging"
	"log"
)

func main() {
	log.Print("config init")
	cfg := config.GetConfig()

	log.Print("logger init")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Application is running")
	a.Run()
}
