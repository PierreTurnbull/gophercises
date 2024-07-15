package main

import (
	"fmt"
	"net/http"
	"urlshort/urlshort"
)

func main() {
	defaultHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	handlerMap := urlshort.GetHandlerMap()

	for key, value := range handlerMap {
		http.HandleFunc(key, value)
	}
	http.HandleFunc("/", defaultHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}
