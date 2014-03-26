package main

import (
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()

	if path != "/" && !strings.HasPrefix(path, "/new") {
		http.Redirect(w, r, "http://old.asturix.com" + path, 301)
	} else {
		log.Println(path)
		if path == "/" {
			path = "/index.html"
		}
		log.Println(path[1:])
		http.ServeFile(w, r, path[1:])
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started")
}