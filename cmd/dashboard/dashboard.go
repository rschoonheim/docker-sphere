package main

import (
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//todo: check authentication.

		// Does the path contains the `/assets/` string?
		if strings.Contains(r.URL.Path, "/assets/") {
			http.ServeFile(w, r, "web/dist/"+r.URL.Path)
		} else {
			http.ServeFile(w, r, "web/dist/index.html")
		}

	})

	http.ListenAndServe(":8080", nil)

	select {}

}
