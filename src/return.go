package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
)

// To add json-iterator
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ReturnIndex is the response while the index url is accessed.
func ReturnIndex(output http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(output, "ImageManagerAPI v1.0")
}

// ReturnAllImages is the response to return all the images
func ReturnAllImages(output http.ResponseWriter, reader *http.Request) {
	dummy := Images{
		Image{
			ID:        0,
			Image:     "Img",
			Thumbnail: "Thumbnail",
			Caption:   "Caption",
		},
		Image{
			ID:        1,
			Image:     "Img",
			Thumbnail: "Thumbnail",
			Caption:   "Caption",
		},
	}

	output.Header().Set("Content-Type", "application/json")
	output.WriteHeader(http.StatusOK)
	json.NewEncoder(output).Encode(dummy)

	fmt.Println("Endpoint Hit: returnAllImages")
}

// ReturnSingleImage This is to return the single image
func ReturnSingleImage(output http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	imageID := vars["id"]
	fmt.Fprintf(output, "imageId: "+imageID)
}
