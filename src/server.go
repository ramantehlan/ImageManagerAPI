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

	router := GetRouter()
	server := &http.Server{
		Addr:           ":" + viper.GetString("server.port"),
		Handler:        router,
		ReadTimeout:    time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout:   time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
		MaxHeaderBytes: viper.GetInt("server.maxHeaderBytes") << 20,
	}

	fmt.Println("ImageManagerAPI v1.0")
	fmt.Println("Running server on: http://0.0.0.0:" + viper.GetString("server.port"))
	log.Fatal(server.ListenAndServe())
}
