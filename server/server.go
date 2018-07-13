package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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
		logger.Println("Running server on " + listenAddress)
		logger.Fatal(srv.ListenAndServe())
	}()

	return srv
}
