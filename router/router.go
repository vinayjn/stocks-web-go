package router

import (
	"net/http"
	"stocks/handlers"
)

// HandleURL ...
func HandleURL() {
	http.HandleFunc("/", handlers.Index) // set router
	http.HandleFunc("/search", handlers.SearchStocks)
	http.HandleFunc("/quote", handlers.QuoteStocks)
}
