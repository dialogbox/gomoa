package main

import (
	"log"

	"github.com/zeromq/goczmq"
)

func main() {
	server, err := goczmq.NewRouter("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Destroy()

	log.Println("server created and bound")

	// Create a server socket and connect it to the router.
	client, err := goczmq.NewDealer("tcp://127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Destroy()

	log.Println("client created and connected")

	for {
		request, err := server.RecvMessage()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("router received '%s' from '%v'", request[1], request[0])
	}
}
