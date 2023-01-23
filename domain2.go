package main

import (
	"log"
	"net/http"
)

// Attacker's domain

func main() {
	log.Fatal(http.ListenAndServe("localhost:8001", http.FileServer(http.Dir("./web"))))
}
