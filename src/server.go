package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	fmt.Println("ImageManagerAPI v1.0") 
	router := GetRouter()	
	log.Fatal(http.ListenAndServe(":8081",router))
}
