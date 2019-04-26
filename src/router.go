package main

import (
	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{"Index", "GET", "/", ReadIndex},
	Route{"Index", "GET", "/v1", ReadIndex},
	Route{"GetAllImages", "GET", "/v1/image/all", ReadAllImages},
	Route{"GetSingleImage", "GET", "/v1/image/{id}", ReadSingleImage},
	Route{"PostSingleImage", "POST", "/v1/image/upload", CreateSingleImage},
}

// GetRouter is to get all the routers which this API will serve
func GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
