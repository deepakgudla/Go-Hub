package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define routes and their corresponding handler functions
	http.HandleFunc("/", homeHandler)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response content type
	w.Header().Set("Content-Type", "text/html")

	// Write the response body
	fmt.Fprintf(w, "<h1> GOLANG SERVER RESPONSE </h1>")
}
