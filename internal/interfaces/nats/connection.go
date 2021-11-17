package nats

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var once sync.Once

type NatsCon struct {
	Con *nats.Conn
}

var (
	instance *NatsCon
)

func GetNatsCon() *NatsCon {
	once.Do(func() {
		//cfg := configuration.GetConfig()
		var err error

		nc, err := nats.Connect(nats.DefaultURL,
			nats.RetryOnFailedConnect(true),
			nats.MaxReconnects(10),
			nats.ReconnectWait(time.Second),
			nats.ReconnectHandler(func(_ *nats.Conn) {
				// Note that this will be invoked for the first asynchronous connect.
			}))
		if err != nil {
			log.Error(err)
		}

		instance = &NatsCon{Con: nc}
	})

	return instance
}
