package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Victim's domain

const enableCORS = false

var token = time.Now().Format("2006-01-02_15:04:05.000000000")

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server with token: " + token)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var hasCookie bool
	for k, v := range r.Header {
		// fmt.Printf("Header field %q, Value %q\n", k, v)
		if k == "Cookie" {
			fmt.Println(strings.Join(v, ""))
			hasCookie = strings.Contains(strings.Join(v, ""), "sessionid="+token)
		}
	}

	fmt.Printf("Got request - %s %s %s hasCookie=%v\n", r.Method, r.URL, r.Proto, hasCookie)

	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Set-Cookie", "sessionid="+token+"; path=/; httponly")
	if enableCORS {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}
	fmt.Fprintf(w, "Got request - %s %s %s hasCookie=%v\n", r.Method, r.URL, r.Proto, hasCookie)
}
