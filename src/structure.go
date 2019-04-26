package main

import (
	"net/http"
)

// Image is the structure of a single image entry
type Image struct {
	ID        int    `json:"id"`
	Image     string `json:"image"`
	Thumbnail string `json:"thumbnail"`
	Caption   string `json:"caption"`
}

// Images is the array of multiple Image
type Images []Image

// Route is the structure of a single route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the array of multiple Route
type Routes []Route

// Dir is to store single important paths for API
type Dir struct {
	path string
}

// Dirs is to stroe multiple paths
type Dirs []Dir
