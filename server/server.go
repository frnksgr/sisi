package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func byteSize(s string) int64 {
	if len(s) > 0 {
		var factor int64
		switch f := s[len(s)-1]; f {
		case 'B':
			factor = 1
		case 'K':
			factor = 1024
		case 'M':
			factor = 1024 * 1024
		case 'G':
			factor = 1024 * 1024 * 1024
		}
		//n, err := strconv.Atoi(s[:len(s)-1])
		n, err := strconv.ParseInt(s[:len(s)-1], 10, 64)
		if err == nil {
			return n * factor
		} else {
			log.Println(err)
		}
	}
	return 0
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hossa!")
}

func chunkedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "To be implemented")
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	defSize := "10M"
	n := byteSize(defSize)

	params := r.URL.Query()

	if keys, ok := params["size"]; ok {
		n = byteSize(keys[0])
	}

	if n == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Wrong Parameter")
	} else {
		input := "/dev/zero"
		if _, ok := params["random"]; ok {
			input = "/dev/urandom"
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("ContentType", "application/octet-stream")
		data, _ := os.Open(input)
		defer data.Close()

		io.CopyN(w, data, n)
	}
}

// RunServer runs web server on given address e.g. ":4711"
func RunServer(listenAddress string) *http.Server {

	// use DefaultMux for now add middleware later
	http.HandleFunc("/", logRequest(helloHandler))
	http.HandleFunc("/data", logRequest(dataHandler))

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
