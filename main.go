package main

import (
	"log"
	"net/http"
	"stocks/router"
)

func main() {
	router.HandleURL()                              // set router
	err := http.ListenAndServe("0.0.0.0:9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
