package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
)

// To add json-iterator
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ReadIndex is the response while the index url is accessed.
func ReadIndex(output http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(output, "ImageManagerAPI v1.0")
}

// ReadAllImages is the response to return all the images
func ReadAllImages(output http.ResponseWriter, reader *http.Request) {
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

	Log("info", "Endpoint Hit: ReadAllImages")
}

// ReadSingleImage This is to return the single image
func ReadSingleImage(output http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	imageID := vars["id"]
	fmt.Fprintf(output, "imageId: "+imageID)

	Log("info", "Endpoint Hit: ReadSingleImage")
}
