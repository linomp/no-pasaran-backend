package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /oauthdemo", func(w http.ResponseWriter, r *http.Request) {
		oauthDemo(w, r)
	})

	mux.HandleFunc("GET /google-oauth-callback", func(w http.ResponseWriter, r *http.Request) {
		data, _ := getUserDataFromGoogle(r.FormValue("code"))
		html, _ := getAsHtml(data)

		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		metrics := getStatus(r)
		html := generateServerStatusHTML(metrics)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	log.Println("Running on http://localhost:8001")
	log.Println("Try oauth demo on http://localhost:8001/oauthdemo")
	log.Fatal(http.ListenAndServe("0.0.0.0:8001", mux))
}
