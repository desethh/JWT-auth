package main

import (
	"os"

	controller "jwt/controllers"
	"jwt/dependencies"

	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func setupController(dependencies dependencies.Dependencies) controller.Controller {
	controller := controller.NewController(dependencies.AuthService)
	return *controller
}
