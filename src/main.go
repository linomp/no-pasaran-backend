package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occurred loading env variables: %s", err)
	}

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

	log.Printf("Running on http://%s:%s\n", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Printf("Try oauth demo on http://%s:%s/oauthdemo\n", os.Getenv("HOST"), os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), mux))
}
