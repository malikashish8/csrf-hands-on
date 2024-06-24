package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Attacker's domain

// get a list of all files in web folder
func allWebFiles() []string {
	files, err := filepath.Glob("web/*")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" || r.URL.Path == "/" {
		output :=(`
<!DOCTYPE html>
<link rel="icon" href="data:;base64,iVBORw0KGgo=">
<h1>Attacker Domain (port 8001)</h1>
First initialize a session on the victim's domain (port 8000) by clicking on init.html. Then try the various link.<br><br>
	`)
		// add links from all web files to the output
		output += `<ul>`
		for _, f := range allWebFiles() {
			output += `<li><a href="/` + f + `">` + strings.Split(f, "- ")[1] + `</a></li>`
		}
		output += `</ul>`
		// send the output to the client
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(output))
	} else {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		r.URL.Path = filepath.Clean(r.URL.Path)
		// ensure that the filepath is within the web folder
		if strings.HasPrefix(r.URL.Path, "/web/") {
			http.ServeFile(w, r, r.URL.Path[1:])
		} else {
			http.NotFound(w, r)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting started - https://localhost:8001")
	log.Fatal(http.ListenAndServeTLS("localhost:8001", "localhost.crt", "localhost.key", nil))
}
