package main

import (
	"github.com/sirupsen/logrus"

	"polling-system/internal/app"
	"polling-system/internal/config"
)

func main() {
	conf, err := config.ParseConfig()
	if err != nil {
		logrus.Fatalf("invalid config %v", err.Error())
	}

	err = app.Run(conf)
	if err != nil {
		logrus.Fatalf("error running app: %v", err)
	}

	logrus.Info("Service successfully stopped")
}
