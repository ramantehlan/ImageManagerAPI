package main

import (
	"net/http"
)

// Image is the structure of a single image entry
type Image struct {
	ID        int    `json:"ID"`
	Image     string `json:"Image"`
	Thumbnail string `json:"Thumbnail"`
	Caption   string `json:"Caption"`
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
