package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/static/"))

	// start HTTP server with `http.DefaultServeMux` handler
	log.Fatal(http.ListenAndServe(":8080", fs))
}
