package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tedoham/mytempl/internal/templates"
)

// Initialize the file server for static content in the init function
func init() {
	// Serve CSS statically
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))
}

func main() {

	// Use a template that doesn't take parameters.
	http.Handle("/", templ.Handler(templates.Hello("yooooooooooo.....")))

	// Use a template that accesses data or handles form posts.
	// http.Handle("/posts", NewPostsHandler())

	// Start the server.
	fmt.Println("listening on http://localhost:3000")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
