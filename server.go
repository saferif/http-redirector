package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Host != "" {
			http.Redirect(w, r, "http://www."+r.Host, http.StatusMovedPermanently)
		} else {
			http.Error(w, "Host header must be specified", http.StatusBadRequest)
		}
	})
	http.ListenAndServe(":80", nil)
}
