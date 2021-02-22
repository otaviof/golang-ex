package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting web-server on port 8080...")
	fs := http.FileServer(http.Dir("assets/static/"))
	err := http.ListenAndServe(":8080", fs)
	log.Fatalf("[ERROR] '%#v'", err)
}
