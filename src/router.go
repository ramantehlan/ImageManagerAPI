package main

import (
	"github.com/gorilla/mux"
)

var routes = Routes{
	Route{"Index","GET","/", ReturnIndex},
	Route{"Index","GET","/v1", ReturnIndex},
	Route{"GetAllImages", "GET", "/v1/image/all", ReturnAllImages},
	Route{"GetSingleImage", "GET", "/v1/image/{id}", ReturnSingleImage},
}

// GetRouter is to get all the routers which this API will serve
func GetRouter() *mux.Router{
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