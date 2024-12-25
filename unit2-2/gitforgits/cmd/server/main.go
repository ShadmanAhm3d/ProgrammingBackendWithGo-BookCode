package main


import (
  "net/http"
  "gitforgits.com/internal/api"
)

func main(){
  http.HandleFunc("/",api.HandleIndex);
  http.HandleFunc("/books",api.HandleBooks);


  http.ListenAndServe(":8080", nil)
}
