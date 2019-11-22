package main

import (
	"log"
	"net/http"
	"stocks/router"
)

func main() {
	router.HandleURL()
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
