package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := os.Getenv("ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(OKHandler))
	mux.Handle("/fake", http.HandlerFunc(NotFoundHandler))

	server := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	log.Printf("Listening on %s", listenAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

// OKHandler responds to web requests with a 200 OK.
func OKHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusText(http.StatusOK), http.StatusOK)
}

// NotFoundHandler responds to web requests with a 404 Not Found.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// Response is the HTTP response sent to clients.
type Response struct {
	Message string `json:"message"`
}

// respond is a helper function for sending a JSON response to a client.
func respond(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Message: message}
	json.NewEncoder(w).Encode(resp)
}
