package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("welcome to publisher's world!")
	publish()
}

// func encodedConnection(nc *nats.Conn) nats.EncodedConn {
// 	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ec.Close()
// 	return *ec
// }

func publish() {

	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Publish the message

	if err := nc.Publish("msg.test", []byte("ALl is well")); err != nil {
		log.Fatal(err)
	}

}
