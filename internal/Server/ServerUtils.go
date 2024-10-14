package server

import (
	"log"
	"net/http"

	"swiftgo/internal/request"
	"swiftgo/internal/response"
	"swiftgo/internal/Router"
)

func StartServerCoroutine(host string, port string) {
	startServerErr := http.ListenAndServe(host + ":" + port, nil)
	if startServerErr != nil {
			log.Fatal(startServerErr)
	}
}

func WrapMiddleware(next router.HandlerType, middleware router.MiddlewareType)router.HandlerType {
	return func(r *request.Request, w response.Response) {
			if !middleware(r, w) {
					return
			}

			next(r, w)
	}
}