package main

import (
	log "github.com/sirupsen/logrus"
	"main/internal/configuration"
	"main/internal/infrastructure/nats"
	"main/internal/interfaces/events/listener"
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
	// add here the nats event listener
	natsConn := nats.GetNatsCon()

	err := listener.UserEventsListner(natsConn.Con)
	if err != nil {
		log.Error(err)
	}

}
