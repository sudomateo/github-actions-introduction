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

	app := App{
		Log: log.New(os.Stdout, "example-app: ", 0),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(app.OKHandler))
	mux.Handle("/fake", http.HandlerFunc(app.NotFoundHandler))

	server := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	log.Printf("Listening on %s", listenAddr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

// App represents an example web application.
type App struct {
	Log *log.Logger
}

// OKHandler responds to web requests with a 200 OK.
func (a App) OKHandler(w http.ResponseWriter, r *http.Request) {
	a.Log.Printf("%s", r.URL.Path)
	respond(w, http.StatusText(http.StatusOK), http.StatusOK)
}

// NotFoundHandler responds to web requests with a 404 Not Found.
func (a App) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	a.Log.Printf("%s", r.URL.Path)
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
