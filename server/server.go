package server

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hossa!")
}

func chunkedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// RunServer runs web server on given address e.g. ":4711"
func RunServer(listenAddress string) *http.Server {

	// use DefaultMux for now add middleware later
	http.HandleFunc("/", logRequest(helloHandler))

	srv := &http.Server{
		Addr: listenAddress,
		//Handler: mux,
		//WriteTimeout: time.Second * 60,
		//ReadTimeout:  time.Second * 60,
		//IdleTimeout:  time.Second * 60,
	}

	go func() {
		log.Println("Running server on " + listenAddress)
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}
