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
userController := func(req *request.Request, res: response.Response) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"name": "John"}`))
}
```

### Response
- Native Http Responses have been extended with named methods for the responses
```go
userController := func(req *request.Request, res: response.Response) {
		res.Ok([]byte(`{"name": "John"}`))
}

userController2 := func(req *request.Request, res: response.Response) {
		res.NoContent()
}
```

### Global Middlewares
- You can define global middlewares that will be executed at the start of the middleware chain on every handler

```go
globalMiddleware := func(req *request.Request, res: response.Response) bool {
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(http.StatusOK)
  return true
}

app.UseGlobalMiddleware(globalMiddleware)
```

### Middlewares
- Middlewares can both stop middleware chain and return to the client or continue to the next function based on the return of the middleware
```go
returningMiddleware := func(req *request.Request, res: response.Response) bool {
  log.Println("Middleware 1")
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(http.StatusOK)
  res.Write([]byte(`{"message": "Hello, World from Middleware!"}`))
  return false // Stops middleware chain
}

nonReturningMiddleware := func(req *request.Request, res: response.Response) bool {
  log.Println("Keep going!")
  return true // Continues middleware chain
}
```

### Router

```go
userRoutes := router.Router{}

// You can define a route with it's own method, a list of middlewares and an handler
userRoutes.Get("/users", []router.MiddlewareType{middleware1}, func(req *request.Request, res: response.Response) {
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(http.StatusOK)
  res.Write([]byte(`{"message": "Hello, World!"}`))
})

// You can also create a group of middlewares with common initial path and middlewares
userRoutes.Group("nested", []router.MiddlewareType{middleware1}, func(router *router.Router) {
  router.Get("/internal", []router.MiddlewareType{middleware2}, func(req *request.Request, res: response.Response) {
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)
    res.Write([]byte(`{"message": "Hello, World!"}`))
  }) // Final route will be nested/internal and will use middleware1 and middleware2 before handler execution
})

// Aside routes definition, to use a specific user you must call
app.UseRouter(&userRoutes)
```