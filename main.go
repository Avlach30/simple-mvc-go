package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("Server connected at http://localhost:5000")

	error := http.ListenAndServe(":5000", mux)

	log.Fatal(error)
}
