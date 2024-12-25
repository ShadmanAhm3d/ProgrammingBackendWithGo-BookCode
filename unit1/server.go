package main

import (
	"fmt"
	"net/http"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		fmt.Fprint(w, " this is some of the books ")

	case "POST":

		fmt.Fprint(w, " this is  Shit posting of the books ")

	default:
		http.Error(w, "Shadman,this is not a correct request", http.StatusMethodNotAllowed)
	}

}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hwllo")

	})

	http.HandleFunc("/books", BooksHandler)

	fmt.Println("Server has started on Port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Some error has occcured")
	}

	return
}
