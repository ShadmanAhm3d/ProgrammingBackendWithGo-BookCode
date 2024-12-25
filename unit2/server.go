package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is from mainHandler")
	fmt.Print(r)
}

func Bookhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is from books handler")
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is from OrderBooks handler")
}

func main() {

	var mux *http.ServeMux = http.NewServeMux()

	//register handler

	mux.HandleFunc("/", mainHandler)

	mux.HandleFunc("/books", Bookhandler)

	mux.HandleFunc("/orders", OrderHandler)

	//server starting
	fmt.Println("Server starting on PORT 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Some error : ", err)
	}

  return;
}
