package main

import (
	"log"
	"net/http"
)

// home is a handler that takes in two parameters
// http.ResponseWriter provides methods for assembling
// an HTTP response and sending it to the user
// ----
// http.Request is a struct which holds information
// about the current request
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	// initializing new servemux, then registering
	// the function as handler for desired url
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000 port")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
