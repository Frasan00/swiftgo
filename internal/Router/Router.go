package router

import (
	"swiftgo/internal/request"
	"swiftgo/internal/response"
)

type MiddlewareType func(*request.Request, response.Response) bool
type HandlerType func(*request.Request, response.Response)

type Route struct {
	Method  string
	Path    string
	Middlewares []MiddlewareType
	Handler HandlerType
}

type Router struct {
	Routes []Route
}

func (router *Router) AddRoute(route Route) {
	isMethodValid := ValidateHttpMethod(route.Method)
	if !isMethodValid {
		panic("Invalid HTTP Method")
	}

	isPathValid, path := ValidatePath(route.Path)
	if !isPathValid {
		panic("Invalid path in route " + path)
	}

	router.Routes = append(router.Routes, route)
}

func (router *Router) Group(prefix string, middlewares []MiddlewareType, cb func(*Router)) {
	isPathValid, prefixPath := ValidatePath(prefix)
	if !isPathValid {
		panic("Invalid path in route " + prefixPath)
	}

	groupRouter := &Router{}

	cb(groupRouter)

	for _, route := range groupRouter.Routes {
			validPath, path := ValidatePath(route.Path)
			if !validPath {
				panic("Invalid path in route " + path)
			}

			route.Path = prefixPath + path
			route.Middlewares = append(middlewares, route.Middlewares...)
			router.Routes = append(router.Routes, route)
	}
}

func (router *Router) Get(path string, middlewares []MiddlewareType, handler HandlerType) {
	isPathValid, path := ValidatePath(path)
	if !isPathValid {
		panic("Invalid path in route " + path)
	}

	router.Routes = append(router.Routes, Route{
			Method:  "GET",
			Path:    path,
			Middlewares: middlewares,
			Handler: handler,
	})
}

func (router *Router) Post(path string, middlewares []MiddlewareType, handler HandlerType) {
	isPathValid, path := ValidatePath(path)
	if !isPathValid {
		panic("Invalid path in route " + path)
	}

	router.Routes = append(router.Routes, Route{
			Method:  "POST",
			Path: 	path,
			Middlewares: middlewares,
			Handler: handler,
	})
}

func (router *Router) Put(path string, middlewares []MiddlewareType, handler HandlerType) {
	isPathValid, path := ValidatePath(path)
	if !isPathValid {
		panic("Invalid path in route " + path)
	}

	router.Routes = append(router.Routes, Route{
			Method:  "PUT",
			Path: 	path,
			Middlewares: middlewares,
			Handler: handler,
	})
}

func (router *Router) Delete(path string, middlewares []MiddlewareType, handler HandlerType) {
	isPathValid, path := ValidatePath(path)
	if !isPathValid {
		panic("Invalid path in route " + path)
	}

	router.Routes = append(router.Routes, Route{
			Method:  "DELETE",
			Path: 	path,
			Middlewares: middlewares,
			Handler: handler,
	})
}
