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
		dump, err := httputil.DumpRequest(r, false)
		if err != nil {
			log.Print(err)
		} else {
			old, new := []byte{'\r', '\n'}, []byte{'\r', '\n', ' ', ' '}
			bytes.Replace(dump, old, new, bytes.Count(dump, old)-1)
			log.Printf("%s", bytes.Replace(dump, old, new, bytes.Count(dump, old)-1))
		}
		next(w, r)
	}
}
