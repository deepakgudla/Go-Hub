package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Typewriter Font Page</title>
			<style>
				body {
					margin: 0;
					background-color: black;
					color: white;
					height: 100vh;
					display: flex;
					justify-content: center;
					align-items: center;
					font-family: "Courier New", Courier, monospace;
					font-size: 2em;
				}
			</style>
		</head>
		<body>
			<div>I look forward to the wisdom you create 🚀</div>
		</body>
		</html>
	`
	fmt.Fprint(w, html)
}

func main() {
	//.env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default port 8080")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
