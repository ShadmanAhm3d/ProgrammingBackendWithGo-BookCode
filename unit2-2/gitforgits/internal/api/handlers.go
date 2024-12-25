package api

import (
	"fmt"
	"net/http"
	"gitforgits.com/internal/model"
)

// HandleIndex serves the main page.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to gitforgits.com!")
}

// HandleBooks displays books available.
func HandleBooks(w http.ResponseWriter, r *http.Request) {
	books := model.FetchBooks() // Assume FetchBooks returns a list of books

	for _, book := range books {
		fmt.Fprintf(w, "Book: %s\n", book.Title)
	}
}
