package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	origin, _ := url.Parse("http://localhost:3000/")

	director := func(req *http.Request) {
		fmt.Printf("%s => %s\n", req.URL.Host, origin.RequestURI())

		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = origin.Scheme
		req.URL.Host = origin.Host
	}

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	fmt.Println("Proxy listening on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
