package main 

import (
	"fmt"
	"strconv"
	"net/http"
	"html/template"
	"log"
)



func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // Include the footer partial in the template files.
    files := []string{
        "./ui/html/home.page.tmpl",
        "./ui/html/base.layout.tmpl",
        "./ui/html/footer.partial.tmpl",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}
func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}


	// We’re able to do pass in w to Fprintf function the io.Writer type is 
	// an interface, and the http.ResponseWriter object satisfies the 
	// interface because it has a w.Write() method.
	fmt.Fprintf(w, "Display a specific snippet with %s", id)
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
