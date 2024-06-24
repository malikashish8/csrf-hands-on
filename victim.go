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
	log.Fatal(http.ListenAndServeTLS("localhost:8000", "localhost.crt", "localhost.key", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var hasCookie bool
	for k, value := range r.Header {
		// fmt.Printf("Header field %q, Value %q\n", k, value)
		if k == "Cookie" {
			hasCookie = strings.Contains(strings.Join(value, ""), "sessionid="+token)
		}
	}

	fmt.Printf("%s %s hasCookie=%v\n", r.Method, r.URL, hasCookie)

	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Set-Cookie", "sessionid="+token+"; path=/; httponly; secure; SameSite=strict")
	if enableCORS {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	}

	message := fmt.Sprintf("%s %s hasCookie=%v\n", r.Method, r.URL, hasCookie)
	fmt.Fprint(w, getPage(message))
}

func getPage(message string) string {
	m := `<!DOCTYPE html><link rel="icon" href="data:;base64,iVBORw0KGgo="><h1>Victim Domain (port 8000)</h1><pre>` + message + `</pre>`
	return m
}
