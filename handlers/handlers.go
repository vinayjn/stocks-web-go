package handlers

import (		
	"strings"
	"net/http"
	"io/ioutil"		
	"stocks/utils"	
	"stocks/models"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
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
	resp_body, _ := ioutil.ReadAll(response.Body)	
	obj := models.SymbolSearchResponse{}	
	jsonpb.Unmarshal(strings.NewReader(string(resp_body)), &obj)
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
	resp_body, _ := ioutil.ReadAll(response.Body)	
	obj := models.StockQuoteResponse{}	
	jsonpb.Unmarshal(strings.NewReader(string(resp_body)), &obj)
	data, _ := proto.Marshal(&obj)
	w.Write(data)	
}
