package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/nats.go"
)

func main() {

	var queueName = flag.String("q", "NATS-RPLY-22", "Queue Group Name")

	//NATS connection
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	args := os.Args

	subj, reply, i := args[1], args[2], 0

	nc.QueueSubscribe(subj, *queueName, func(msg *nats.Msg) {
		i++
		log.Printf("[#%d] Received on [%s]: '%s'\n", i, msg.Subject, string(msg.Data))
		msg.Respond([]byte(reply))
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subj)

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatalf("Exiting")

}
