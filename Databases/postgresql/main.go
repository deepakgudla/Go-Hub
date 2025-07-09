package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date,omitempty"`
	ISBN          string `json:"isbn,omitempty"`
}

var db *pgxpool.Pool

func main() {
	dbURL := "postgres://user:password@localhost:5432/bookdb"
	if os.Getenv("DB_URL") != "" {
		dbURL = os.Getenv("DB_URL")
	}

	var err error
	db, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Router setup
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Server configuration
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

// Handlers
func getBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(context.Background(), "SELECT id, title, author, published_date, isbn FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		var pubDate time.Time
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &pubDate, &b.ISBN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b.PublishedDate = pubDate.Format("2006-01-02")
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var b Book
	var pubDate time.Time
	err = db.QueryRow(context.Background(),
		"SELECT id, title, author, published_date, isbn FROM books WHERE id = $1", id).Scan(
		&b.ID, &b.Title, &b.Author, &pubDate, &b.ISBN)

	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	b.PublishedDate = pubDate.Format("2006-01-02")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert published_date string to time.Time
	pubDate, err := time.Parse("2006-01-02", b.PublishedDate)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	var id int
	err = db.QueryRow(context.Background(),
		"INSERT INTO books (title, author, published_date, isbn) VALUES ($1, $2, $3, $4) RETURNING id",
		b.Title, b.Author, pubDate, b.ISBN).Scan(&id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pubDate, err := time.Parse("2006-01-02", b.PublishedDate)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(context.Background(),
		"UPDATE books SET title=$1, author=$2, published_date=$3, isbn=$4 WHERE id=$5",
		b.Title, b.Author, pubDate, b.ISBN, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	b.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(context.Background(), "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
