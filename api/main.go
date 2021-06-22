package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	http.Handle("/", r)
	log.Println("Server lisening to port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
