package client

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
)

// RunClient run client against server
func RunClient(serverAddress string, cmd string) {
	log.Println("Running client against " + serverAddress)

	resp, err := http.Get("http://" + serverAddress + "/" + "command")
	if err != nil {
		log.Fatal(err)
	}
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		log.Fatal(err)
	}
	old, new := []byte{'\r', '\n'}, []byte{'\r', '\n', ' ', ' '}
	bytes.Replace(dump, old, new, bytes.Count(dump, old)-1)
	log.Printf("%s", bytes.Replace(dump, old, new, bytes.Count(dump, old)-1))
}
