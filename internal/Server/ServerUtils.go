package server

import (
	"log"
	"net/http"
)

func StartServerCoroutine(host string, port string) {
	startServerErr := http.ListenAndServe(host + ":" + port, nil)
	if startServerErr != nil {
			log.Fatal(startServerErr)
	}
}

func WrapMiddleware(next http.HandlerFunc, middleware func(http.ResponseWriter, *http.Request) bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			if !middleware(w, r) {
					return
			}

			next(w, r)
	}
}