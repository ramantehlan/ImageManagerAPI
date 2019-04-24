package main

import (
	"fmt"
	"net/http"
	"github.com/json-iterator/go"
	"github.com/gorilla/mux"
)

// To add json-iterator
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ReturnIndex(output http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(output, "ImageManagerAPI v1.0")
}

func ReturnAllImages(output http.ResponseWriter, reader *http.Request){
	dummy := Images{
		Image{
			Id : 0,
			Image: "Img",
			Thumbnail: "Thumbnail",
			Caption: "Caption",
		},
		Image{
			Id : 1,
			Image: "Img",
			Thumbnail: "Thumbnail",
			Caption: "Caption",
		},
	}
	
	output.Header().Set("Content-Type", "application/json")
	output.WriteHeader(http.StatusOK)
	json.NewEncoder(output).Encode(dummy)

	fmt.Println("Endpoint Hit: returnAllImages")
}

func ReturnSingleImage(output http.ResponseWriter, reader *http.Request){
	vars := mux.Vars(reader)
	imageId := vars["id"]
	fmt.Fprintf(output, "imageId: " + imageId)
}
