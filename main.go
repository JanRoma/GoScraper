package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This application is very cool.\n")
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/about", about)
	return r
}

func main() {
	router := router()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}