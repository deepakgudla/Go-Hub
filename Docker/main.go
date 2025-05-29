package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Jai Mahishmathi ✊</h1>")
}

func getPort() (string, error) {
	_ = godotenv.Load() // Optional: log warning if needed

	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("PORT not set in environment")
	}
	return port, nil
}

func main() {
	port, err := getPort()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	http.HandleFunc("/", handler)
	fmt.Printf("Server starting at :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
