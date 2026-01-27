package main

import (
	"log"
	"net/http"

	"github.com/nitin0409sep/students-api/internal/config"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()

	if(err != nil) {
		log.Fatal("Failed to start server")
	}

}
