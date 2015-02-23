package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Host != "" {
			http.Redirect(w, r, "http://www."+r.Host, 301)
		} else {
			w.WriteHeader(400)
		}
	})
	http.ListenAndServe(":80", nil)
}
