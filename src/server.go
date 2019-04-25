package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func main() {

	// Load Config file
	LoadConfig()
	// To start the logger and define the default things
	LoadLogger()

	// Fetch routes
	router := GetRouter()
	// Configure the server
	server := &http.Server{
		Addr:           ":" + viper.GetString("server.port"),
		Handler:        router,
		ReadTimeout:    time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
		MaxHeaderBytes: viper.GetInt("server.maxHeaderBytes") << 20,
	}

	fmt.Println(viper.GetString("name"))
	fmt.Println("Running server on: http://0.0.0.0:" + viper.GetString("server.port"))
	fmt.Print("\n--LOG--\n\n")
	Log("info", "Server started at port:"+viper.GetString("server.port"))
	// Connect to DB
	ConnectDB()
	log.Fatal(server.ListenAndServe())
}
