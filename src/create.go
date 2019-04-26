package main

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/spf13/viper"
)

// CreateSingleImage is the response to POST when a image is uploaded
func CreateSingleImage(output http.ResponseWriter, reader *http.Request) {
	Log("info", "Endpoint Hit: CreateSingleImage")
	var img Image

	file, handle, err := reader.FormFile("image")
	ErrorHandler(err)
	defer file.Close()

	switch handle.Header.Get("Content-Type") {
	case "image/jpeg", "image/png":
		saveFile(output, handle, file)
	default:
		output.Header().Set("Content-Type", "application/json")
		output.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(output, "The format file is not valid.")
	}

	img.Caption = reader.FormValue("caption")
	img.Image = viper.GetString("imgAddr") + handle.Filename
	img.Thumbnail = viper.GetString("imgAddr") + handle.Filename
	_, err = DB.Query("INSERT INTO images VALUES(null,?,?,?)", img.Image, img.Thumbnail, img.Caption)
	ErrorHandler(err)

	output.Header().Set("Content-Type", "application/json")
	output.WriteHeader(http.StatusOK)
	JSON.NewEncoder(output).Encode(img)
}

func saveFile(w http.ResponseWriter, handle *multipart.FileHeader, file multipart.File) {
	Log("trace", "Endpoint Hit: saveFile")
	data, err := ioutil.ReadAll(file)
	ErrorHandler(err)
	err = ioutil.WriteFile(viper.GetString("imgAddr")+handle.Filename, data, 0666)
	ErrorHandler(err)
}
