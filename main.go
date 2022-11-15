package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return 
	}

	w.Write([]byte("Hello World"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is the page to show snippets"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		// WriteHeader is only possible to call once per response
		// the default header status code is always 200 OK
		w.Header().Set("Allow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("page to create snippets"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000 port")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
