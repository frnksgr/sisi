package client

import (
	"fmt"
	"log"
)

func hello(serverAddress string) {
	fmt.Println("Hello server on" + serverAddress)
}

// RunClient run client against server
func RunClient(serverAddress string) {
	log.Println("Running client against " + serverAddress)
	hello(serverAddress)
}
