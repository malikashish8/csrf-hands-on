package main

import (
	"log"
	"net/http"
	"os"
)

// Attacker's domain

func main() {

	// http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", http.FileServer(http.Dir("./web"))))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", "application/json")
	// read file from disk

	contents, _ := os.ReadFile("./webpage2.html")
	// write file to response
	w.Write(contents)
}
