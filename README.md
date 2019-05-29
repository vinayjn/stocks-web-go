# Stocks Go

This repo contains the backend code for an iOS app to learn [GoLang](https://golang.org/) and use [Protocol Buffers](https://developers.google.com/protocol-buffers/) as a replacement of JSON based APIs. 



### Setup

I am using [Alpha Vantage]() to fetch the stocks related data. They provide a free to use rate limited API for projects like these. To get this server working you need the following keys in your environment:

```shell
STOCKS_API_KEY="######"
STOCKS_API_BASE_URL="https://www.alphavantage.co/query"
```

Clone this repo in `$GOPATH/src` directory and execute the script `start.sh` script. If everything is setup correctly you can open `http:localhost:9090/`. Check the response message for health check.

You can find the `.proto` files in the [iOS app repo](https://github.com/vinayjn/Stocks-iOS/tree/master/proto/models). 