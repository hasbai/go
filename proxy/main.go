package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func UniversalReverseProxy() *httputil.ReverseProxy {
	director := func(request *http.Request) {
		proxyUrl := "https://" + request.URL.RequestURI()[1:] // RequestURI(): /example.org/path?foo=bar
		request.URL, _ = url.Parse(proxyUrl)
		request.Host = request.URL.Host
		fmt.Println(request.URL)
	}
	return &httputil.ReverseProxy{
		Director: director,
	}
}

func main() {
	proxy := UniversalReverseProxy()
	err := http.ListenAndServe(":10086", proxy)
	if err != nil {
		fmt.Println(err)
	}
}
