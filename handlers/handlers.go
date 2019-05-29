package handlers

import (	
	"fmt"
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
	w.Write([]byte("Hello")) // send data to client side
}

// SearchStocks searches for stocks which matches the query string
func SearchStocks(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	response, _ := utils.MakeGETCall("SYMBOL_SEARCH", map[string]string{"keywords": strings.Join(r.Form["keyword"], "")})
	defer response.Body.Close()	
	resp_body, _ := ioutil.ReadAll(response.Body)	
	obj := models.SymbolSearchResponse{}	
	jsonpb.Unmarshal(strings.NewReader(string(resp_body)), &obj)
	data, _ := proto.Marshal(&obj)
	k, _ := w.Write(data)	
}
