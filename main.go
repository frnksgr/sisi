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

	switch cmd := flag.Arg(0); cmd {
	case "hello":
		client.RunClient(serverAddress, cmd)
	default:
		log.Fatal("unknown Command\n")
	}
}

func main() {
	var serverMode bool
	flag.BoolVar(&serverMode, "s", false,
		"Run server if set, else run client")
	flag.Parse()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	bindAddress := ":" + port
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
