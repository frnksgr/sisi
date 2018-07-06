package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/frnksgr/sisi/client"
	"github.com/frnksgr/sisi/server"
)

func runServer(listenAddress string) {
	srv := server.RunServer(listenAddress)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")
	srv.Shutdown(context.Background())
}

func runClient(serverAddress string) {
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan

		log.Println("Shutdown signal received, exiting...")
		os.Exit(1)
	}()

	client.RunClient(serverAddress)
}

func main() {
	var serverMode bool
	flag.BoolVar(&serverMode, "s", false,
		"Run server if set else run client")

	flag.Parse()

	bindAddress := ":8080"
	serverAddress := os.Getenv("SISI_SERVER")
	if serverAddress == "" {
		serverAddress = "localhost" + bindAddress
	}

	if serverMode {
		runServer(bindAddress)
	} else {
		runClient(serverAddress)
	}
}
