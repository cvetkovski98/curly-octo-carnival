package main

import (
	"flag"
	"log"

	"github.com/cvetkovski98/factory-shared/schemas"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

var (
	publicSubject  string
	privateSubject string
	cluster        string
	service        string
)

func init() {
	flag.StringVar(&service, "service", "node", "The service name of the factory node")
	flag.StringVar(&publicSubject, "publicSubject", "msg.hello", "The publicSubject the node is publishing to")
	flag.StringVar(&privateSubject, "privateSubject", "msg.hello", "The privateSubject the node is publishing to")
	flag.StringVar(&cluster, "cluster", "nats://localhost:4222", "The NATS cluster url the node is connected to")
	flag.Parse()
}

func main() {
	nc, _ := nats.Connect(cluster)
	defer nc.Close()
	log.Printf(
		"[%s] Publishing... "+
			"\nServer: %s "+
			"\nSubjects: %s, %s\n",
		service, cluster, publicSubject, privateSubject,
	)
	publicMessagePublisher := Publisher{Subject: publicSubject}
	privateMessagePublisher := Publisher{Subject: privateSubject}
	publicResponse := &schemas.GetInventoryResponse{Lines: ProductionLines, Type: "public"}
	publicPayload, err := proto.Marshal(publicResponse)
	if err != nil {
		log.Fatalf("[%s] There was an error marshaling data: %s", service, err)
	}
	privateResponse := &schemas.GetInventoryResponse{Lines: ProductionLines, Type: "private"}
	privatePayload, err := proto.Marshal(privateResponse)
	if err != nil {
		log.Fatalf("[%s] There was an error marshaling data: %s", service, err)
	}
	go privateMessagePublisher.StartResponding(nc, privatePayload)
	go publicMessagePublisher.StartResponding(nc, publicPayload)
	select {}
}
