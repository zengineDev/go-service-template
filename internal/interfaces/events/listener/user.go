package listener

import "github.com/nats-io/nats.go"

func UserEventsListner(nc *nats.Conn) error {
	_, err := nc.Subscribe("company.users.*", func(msg *nats.Msg) {

		// parse the cloudevent and call a use case

	})

	if err != nil {
		return err
	}

	return nil
}
