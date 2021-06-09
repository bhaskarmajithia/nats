package main

import (
	"fmt"
	"log"
	"sync"

	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("welcome to subscriber's world!")
	fmt.Println("choose 1 for Synch, 2 for Asynch, and 3 for Queue Sub")
	var choice string
	fmt.Scanln(&choice)
	if choice == "1" {
		subscribeSync()
	} else if choice == "2" {
		subscribeAsync()
	} else if choice == "3" {
		subscribeQueue()
	}

}

func subscribeSync() {

	fmt.Println("Calling subscribeSync ")
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe the message
	sub, err := nc.SubscribeSync("msg.test")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message
	msg, err := sub.NextMsg(100 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Use the response
	log.Printf("Reply: %s", msg.Data)

	if err := sub.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}

func subscribeAsync() {

	fmt.Println("Calling subscribeAsync ")
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(5)

	// Subscribe
	sub, err := nc.Subscribe("msg.test", func(m *nats.Msg) {
		log.Printf("Reply: %s", m.Data)
		wg.Done()
	})

	if err != nil {
		log.Fatal(err)
	}
	// Wait for a message to come in
	wg.Wait()

	//unsubscribe
	if err := sub.Unsubscribe(); err != nil {
		log.Fatal(err)
	}
}

func subscribeQueue() {
	fmt.Println("Calling subscribeQueue ")
	nc, err := nats.Connect("demo.nats.io")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(5)

	//Subscribe to queue with subject
	if _, err := nc.QueueSubscribe("msg.test", "text", func(m *nats.Msg) {
		log.Printf("Reply: %s", m.Data)
		wg.Done()
	}); err != nil {
		log.Fatal()
	}

	wg.Wait()
}
