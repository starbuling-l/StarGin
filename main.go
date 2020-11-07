package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}


// handler echo r.URL.Path
func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}

// handler echo r.URL.Header
func helloHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
}
