package main

import (
	log "github.com/sirupsen/logrus"
	"main/internal/configuration"
	"main/internal/interfaces/cli/commands"
)

func init() {
	cfg := configuration.GetConfig()
	if cfg.App.Environment == configuration.Production {
		log.SetLevel(log.InfoLevel)
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetLevel(log.DebugLevel)
		log.SetFormatter(&log.TextFormatter{})
	}
}

func main() {
	commands.Execute()
}
