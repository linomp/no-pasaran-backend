// Source: https://dev.to/gcdcoder/server-sent-events-in-go-an-efficient-real-time-communication-alternative-4obc

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		for {
			// Simulate sending events every second
			fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.Stamp))
			w.(http.Flusher).Flush()
			time.Sleep(1 * time.Second)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// TODO: turn this into a controller route and send back the template while on another thread (routine? start sending the events)
