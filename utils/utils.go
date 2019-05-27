package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func MakeGETCall(endpoint string, params map[string]string) {
	apiKey, _ := os.LookupEnv("STOCKS_API_KEY")
	apiURL, _ := os.LookupEnv("STOCKS_API_BASE_URL")

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Add("Accept", "application/json")

	// or you can create new url.Values struct and encode that like so
	q := url.Values{}
	q.Add("apikey", apiKey)
	q.Add("function", endpoint)

	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

}
