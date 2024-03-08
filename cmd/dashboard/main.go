package main

import (
	"logging"
	"net/http"
	"strings"
)

func main() {

	logging.Initialize("dashboard")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.Path, "/assets/") {
			http.ServeFile(w, r, "web/dist/"+r.URL.Path)
		} else {
			http.ServeFile(w, r, "web/dist/index.html")
		}

	})

	http.ListenAndServe(":8080", nil)

	select {}

}
