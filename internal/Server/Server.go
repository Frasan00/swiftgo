package server

import (
	"sync"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"swiftgo/internal/Router"
)

type StartServerCb func()

type Server struct {
	Port string
	Host string
	GlobalMiddlewares []func(http.ResponseWriter, *http.Request) bool
}

func (srv* Server) Start(cb StartServerCb) {
	godotenv.Load()

	if srv.Port == "" {
		envPort := os.Getenv("PORT")
		if envPort == "" {
			srv.Port = "80"
		}

		srv.Port = envPort
	}

	if srv.Host == "" {
		envHost := os.Getenv("HOST")
		if envHost == "" {
			srv.Host = "0.0.0.0"
		}

		srv.Host = envHost
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Start the server in an "async" way
	go StartServerCoroutine(srv.Host, srv.Port)

	cb()

	// Infinite loop to keep the server running in background
	wg.Wait()
}

// UseRouter is a function that takes a router and sets up the routes in it in the Server
func (srv *Server) UseRouter(router *router.Router) {
	if srv.GlobalMiddlewares == nil {
			srv.GlobalMiddlewares = []func(http.ResponseWriter, *http.Request) bool{}
	}

	for _, route := range router.Routes {
			finalHandler := route.Handler

			// Apply route-specific middlewares first
			for i := len(route.Middlewares) - 1; i >= 0; i-- {
					middleware := route.Middlewares[i]
					finalHandler = WrapMiddleware(finalHandler, middleware)
			}

			// Apply global middlewares
			for i := len(srv.GlobalMiddlewares) - 1; i >= 0; i-- {
					middleware := srv.GlobalMiddlewares[i]
					finalHandler = WrapMiddleware(finalHandler, middleware)
			}

			http.HandleFunc(route.Path, finalHandler)
	}
}

// Defines a global middleware for the server that will be applied to all routes
func (srv *Server) UseGlobalMiddleware(middleware func(http.ResponseWriter, *http.Request) bool) {
	if srv.GlobalMiddlewares == nil {
			srv.GlobalMiddlewares = []func(http.ResponseWriter, *http.Request) bool{}
	}

	srv.GlobalMiddlewares = append(srv.GlobalMiddlewares, middleware)
}