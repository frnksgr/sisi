package server

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
)

//type middleware func(next http.HandlerFunc) http.HandlerFunc

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d, e := httputil.DumpRequest(r, false)
		if e != nil {
			log.Println(e)
		} else {
			o, n := []byte{'\r', '\n'}, []byte{'\r', '\n', ' ', ' '}
			bytes.Replace(d, o, n, bytes.Count(d, o)-1)
			log.Printf("%s", bytes.Replace(d, o, n, bytes.Count(d, o)-1))
		}
		next(w, r)
	}
}
