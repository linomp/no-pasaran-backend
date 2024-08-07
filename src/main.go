package main

import (
	"log"
	"net/http"
)

type metricsHandler struct {
}

func (h *metricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	metrics := getStatus(r)
	html := generateServerStatusHTML(metrics)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", &metricsHandler{})

	log.Fatal(http.ListenAndServe("0.0.0.0:8001", mux))
}
