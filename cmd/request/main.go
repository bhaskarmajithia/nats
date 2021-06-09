package main

import (
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	log.Println("Sending the request")

	args := os.Args
	subj, data := args[1], []byte(args[2])
	// Send the request

	msg, err := nc.Request(subj, data, 2*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Published [%s] : '%s'", subj, data)
	// Use the response
	log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))

	// Close the connection
	nc.Close()
}
