# swiftgo

Backend Framework written in Go.

## Overview

This project is a backend framework written in Go. It uses environment variables for configuration, sets up a server, and defines routes with middleware support.


## Installation

1. Clone the repository:
```sh
git clone https://github.com/frasan00/swiftgo.git
cd swiftgo
```

2. Install dependencies:
```sh
go mod tidy
```

## Usage

### Initialize a server

```go
// You can omit params in Server creation, the library will check first for envs HOST and PORT, then will give defaults PORT: 80 and HOST: 0.0.0.0
app := server.Server{}

// You can also specify port and host
app := server.Server{
  Port: "80"
  Host: "0.0.0.0"
}
```

### Handlers

```go
userController := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "John"}`))
}
```

### Global Middlewares
- You can define global middlewares that will be executed at the start of the middleware chain on every handler

```go
globalMiddleware := func(w http.ResponseWriter, r *http.Request) bool {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  return true
}

app.UseGlobalMiddleware(globalMiddleware)
```

### Middlewares
- Middlewares can both stop middleware chain and return to the client or continue to the next function based on the return of the middleware
```go
returningMiddleware := func(w http.ResponseWriter, r *http.Request) bool {
  log.Println("Middleware 1")
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "Hello, World from Middleware!"}`))
  return false // Stops middleware chain
}

nonReturningMiddleware := func(w http.ResponseWriter, r *http.Request) bool {
  log.Println("Keep going!")
  return true // Continues middleware chain
}
```

### Router

```go
userRoutes := router.Router{}

// You can define a route with it's own method, a list of middlewares and an handler
userRoutes.Get("/users", []func(http.ResponseWriter, *http.Request) bool{middleware1}, func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`{"message": "Hello, World!"}`))
})

// You can also create a group of middlewares with common initial path and middlewares
userRoutes.Group("nested", []func(http.ResponseWriter, *http.Request) bool{middleware1}, func(router *router.Router) {
  router.Get("/internal", []func(http.ResponseWriter, *http.Request) bool{middleware2}, func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Hello, World!"}`))
  }) // Final route will be nested/internal and will use middleware1 and middleware2 before handler execution
})

// Aside routes definition, to use a specific user you must call
app.UseRouter(&userRoutes)
```