package main

import (
	"net/http"
)

func GenerateServerStatusHTML(r *http.Request) string {
	metrics := getStatus(r)
	return generateServerStatusHTML(metrics)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := GenerateServerStatusHTML(r)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	http.ListenAndServe(":8001", nil)
}
