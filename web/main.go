package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	r.Use(mux.CORSMethodMiddleware(r))

	log.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
