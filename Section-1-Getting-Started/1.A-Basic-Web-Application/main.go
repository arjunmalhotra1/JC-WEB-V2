package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-type", "text/plain")
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1> Welcome to my awesome site </h1>")
}

func contactHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1> Contact Page </h1> <p>To get in touch, email me at <a href=\"mailto:arjun1malhotra@gmail.com\">arjun1malhotra@gmail.com</a>.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on :3000....")
	http.ListenAndServe(":3000", nil)
}
