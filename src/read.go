package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ReadIndex is the response while the index url is accessed.
func ReadIndex(output http.ResponseWriter, reader *http.Request) {
	fmt.Fprintln(output, "ImageManagerAPI v1.0")
	Log("info", "Endpoint Hit: ReadIndex")
}

// ReadAllImages is the response to return all the images
func ReadAllImages(output http.ResponseWriter, reader *http.Request) {
	Log("info", "Endpoint Hit: ReadAllImages")
	results, err := DB.Query("SELECT * FROM images")
	ErrorHandler(err)

	var images Images
	for results.Next() {
		var img Image
		// for each row, scan the result into our image composite object
		err = results.Scan(&img.ID, &img.Image, &img.Thumbnail, &img.Caption)
		ErrorHandler(err)

		// Append images to array
		images = append(images, img)
	}

	defer results.Close()
	output.Header().Set("Content-Type", "application/json")
	output.WriteHeader(http.StatusOK)
	JSON.NewEncoder(output).Encode(images)
}

// ReadSingleImage This is to return the single image
func ReadSingleImage(output http.ResponseWriter, reader *http.Request) {
	Log("info", "Endpoint Hit: ReadSingleImage")
	vars := mux.Vars(reader)
	var img Image
	err := DB.QueryRow("SELECT * FROM images where id = ?", vars["id"]).Scan(&img.ID, &img.Image, &img.Thumbnail, &img.Caption)
	ErrorHandler(err)

	output.Header().Set("Content-Type", "application/json")
	output.WriteHeader(http.StatusOK)
	JSON.NewEncoder(output).Encode(img)
}
