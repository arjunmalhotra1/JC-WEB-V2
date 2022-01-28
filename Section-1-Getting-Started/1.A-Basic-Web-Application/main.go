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

func pathHandler(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		homeHandler(w, req)
	case "/contact":
		contactHandler(w, req)

	default:
		http.Error(w, "Page not found", http.StatusNotFound)
		// This below displays "Not Found"
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "Page Not found")

	}

}

func main() {
	http.HandleFunc("/", pathHandler)
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting server on :3000....")
	http.ListenAndServe(":3000", nil)
}
