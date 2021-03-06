package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is home page\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":9090", nil)
}
