package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	Subject string
}

func (publisher *Publisher) StartResponding(connection *nats.Conn, message []byte) {
	_, err := connection.Subscribe(publisher.Subject, func(msg *nats.Msg) {
		log.Printf("[%s] Received request\n", service)
		err := msg.Respond(message)
		if err != nil {
			log.Fatalf("[%s] There was an error sending response: %s", service, err)
		}
	})
	if err != nil {
		log.Fatalf("[%s] There was an error making a subscription: %s", service, err)
	}
}
