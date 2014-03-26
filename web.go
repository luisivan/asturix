package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path != "/" && !strings.HasPrefix(path, "/new") {
		http.Redirect(w, r, "http://old.asturix.com" + path, 301)
	} else {
		u, err := url.Parse("http://web.asturix.com")
		if err != nil {
			log.Fatal(err)
		}
		 
		reverse_proxy := httputil.NewSingleHostReverseProxy(u)
		reverse_proxy.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started")
}