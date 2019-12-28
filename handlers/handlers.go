package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"stocks/models"
	"stocks/utils"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// Index says hello
func Index(w http.ResponseWriter, r *http.Request) {
	_, message := utils.CheckSetup()
	w.Write([]byte(message)) // send data to client side
}

// SearchStocks searches for stocks which matches the query string
func SearchStocks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response, error := utils.MakeGETCall("SYMBOL_SEARCH", map[string]string{"keywords": strings.Join(r.Form["keyword"], "")})
	if error != nil {
		w.Write([]byte(string(error.Error())))
		return
	}
	defer response.Body.Close()
	respBody, _ := ioutil.ReadAll(response.Body)
	obj := models.SymbolSearchResponse{}
	fmt.Print(string(respBody))
	jsonpb.Unmarshal(strings.NewReader(string(respBody)), &obj)
	data, _ := proto.Marshal(&obj)
	w.Write(data)
}

// QuoteStocks find quotes for stocks
func QuoteStocks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response, error := utils.MakeGETCall("GLOBAL_QUOTE", map[string]string{"symbol": strings.Join(r.Form["symbol"], "")})
	if error != nil {
		w.Write([]byte(string(error.Error())))
		return
	}
	defer response.Body.Close()
	respBody, _ := ioutil.ReadAll(response.Body)
	obj := models.StockQuoteResponse{}
	fmt.Print(string(respBody))
	jsonpb.Unmarshal(strings.NewReader(string(respBody)), &obj)
	data, _ := proto.Marshal(&obj)
	w.Write(data)
}
