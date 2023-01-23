package main

import (
	"fmt"
	"log"
	"net/http"
)

// Victim's domain

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Printf("Got request - %s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	// for k, v := range r.Header {
	// 	fmt.Printf("Header field %q, Value %q\n", k, v)
	// }

	fmt.Printf("Host = %q\n", r.Host)
	fmt.Printf("RemoteAddr= %q\n", r.RemoteAddr)
	fmt.Println("***************************")
	fmt.Fprintf(w, "Got request - %s %s %s \n", r.Method, r.URL, r.Proto)
}
