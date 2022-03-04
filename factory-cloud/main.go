package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cvetkovski98/factory-shared/schemas"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

var (
	subject string
	cluster string
)

func init() {
	flag.StringVar(&subject, "subject", "msg.hello", "The subject the node is publishing to")
	flag.StringVar(&cluster, "cluster", "nats://localhost:4222", "The NATS cluster url the node is connected to")
	flag.Parse()
}

func main() {
	nc, _ := nats.Connect(cluster)
	defer nc.Close()

	fmt.Printf("Listening... \nServer: %s \nSubject: %s\n", cluster, subject)
	http.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		request := schemas.GetInventoryRequest{
			Detail: "Inventory listing request",
		}
		reqJson, err := json.Marshal(request)
		if err != nil {
			fmt.Fprintf(w, "Oops, there was an error: %q", err)
			return
		}
		message, err := nc.Request(subject, reqJson, 10*time.Second)
		if err != nil {
			fmt.Fprintf(w, "Oops, there was a request error: %q", err)
			return
		}
		var inventoryResponse = &schemas.GetInventoryResponse{}
		err = proto.Unmarshal(message.Data, inventoryResponse)
		if err != nil {
			fmt.Fprintf(w, "Oops, there was an unmarshaling error: %q", err)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		json.NewEncoder(w).Encode(inventoryResponse)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
