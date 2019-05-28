package main

import (
	"log"
	"net/http"
	"os"
)

type serv struct{}

func (serv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Handler(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Starting sigil on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, serv{}))
}
