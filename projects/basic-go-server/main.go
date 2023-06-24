package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "./static" directory
	// Any request to the root URL ("/") will be handled by the FileServer
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Start the server and listen on port 8080
	// The server will use the default ServeMux and handle requests using the registered handlers
	log.Fatal(http.ListenAndServe(":8080", nil))
}
