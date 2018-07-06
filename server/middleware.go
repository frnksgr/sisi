package server

import (
	"log"
	"net/http"
	"net/http/httputil"
)

//type middleware func(next http.HandlerFunc) http.HandlerFunc

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequestOut(r, false)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%q", dump)
		}
		next(w, r)
	}
}
