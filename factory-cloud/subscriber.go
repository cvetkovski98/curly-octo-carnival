package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

type Subscriber struct {
	Subject string
}

func (subscriber Subscriber) StartListening(connection *nats.Conn) {
	_, err := connection.Subscribe(subject, func(msg *nats.Msg) {
		log.Println(fmt.Sprintf("[%s] \t %s", msg.Subject, msg.Data))
	})
	if err != nil {
		log.Fatalf("There was an error publishing message: %s", err)
	}
}
