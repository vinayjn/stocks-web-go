package handlers

import (
	"net/http"
	"stocks/utils"
	"strings"
)

// Index says hello
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello")) // send data to client side
}

// SearchStocks searches for stocks which matches the query string
func SearchStocks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.MakeGETCall("SYMBOL_SEARCH", map[string]string{"keywords": strings.Join(r.Form["keyword"], "")})
}
