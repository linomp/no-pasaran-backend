package main

import (
	"log"
	"net/http"
)

type healthcheckHandler struct{}

func (h *healthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/plain")

	switch {
	case r.Method == http.MethodGet:
		msg := []byte("hello there")
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
		return
	default:
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &healthcheckHandler{})

	log.Fatal(http.ListenAndServe("0.0.0.0:8001", mux))
}
