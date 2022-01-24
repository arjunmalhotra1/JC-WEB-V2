package main

import (
	"fmt"
	"net/http"
)

func HandlerFunc(w http.Respo nseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1> Welcome to my awesome site </h1>")
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	fmt.Println("Starting server on :3000....")
	http.ListenAndServe(":3000", nil)
}
