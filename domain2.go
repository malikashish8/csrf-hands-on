package main

import (
	"log"
	"net/http"
)

// Attacker's domain

func main() {
	log.Println("Starting started - https://localhost:8001")
	log.Fatal(http.ListenAndServeTLS("localhost:8001", "localhost.crt", "localhost.key", http.FileServer(http.Dir("./web"))))
}
