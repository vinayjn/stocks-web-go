package utils

import (
	"fmt"	
	"log"
	"net/http"
	"net/url"
	"os"	
	"errors"
)

func CheckSetup() (bool, string) {
	_, apiKeyPresent := os.LookupEnv("STOCKS_API_KEY")	
	if !apiKeyPresent {
		return false, "STOCKS_API_KEY is not set"
	}
	_, apiURLPresent := os.LookupEnv("STOCKS_API_BASE_URL")
	if !apiURLPresent {
		return false, "STOCKS_API_BASE_URL is not set"
	}
	return true, "Success"
}

func MakeGETCall(endpoint string, params map[string]string) (*http.Response, error) {	
	success, message := CheckSetup()	
	if !success {		
		return nil, errors.New(message)
	}
	apiKey, _ := os.LookupEnv("STOCKS_API_KEY")
	apiURL, _ := os.LookupEnv("STOCKS_API_BASE_URL")

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	
	q := url.Values{}
	q.Add("apikey", apiKey)
	q.Add("function", endpoint)

	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()	
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return nil, err
	}
	return resp, nil
}
