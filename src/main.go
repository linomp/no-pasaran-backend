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
		if err != nil {
			log.Println(err.Error())
			return
		}

		html, err := getAsHtml(data)
		if err != nil {
			log.Println(err.Error())
			return
		}

		w.Header().Set("Content-Type", "text/html")

		_, _ = w.Write([]byte(html))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		metrics := getStatus(r)
		html, err := generateServerStatusHTML(metrics)
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(html))
	})

	log.Printf("Running on %s\n", os.Getenv("BASE_URL"))
	log.Printf("Try oauth demo on %s/oauthdemo\n", os.Getenv("BASE_URL"))

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")), mux))
}
