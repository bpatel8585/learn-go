package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Printf("this is a test")

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Printf("error %+v", err)
		return
	}

	fmt.Printf("connected to NATS: %+v", nc)

}
