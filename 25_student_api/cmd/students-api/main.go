package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/nitin0409sep/students-api/internal/config"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	// setup router
	router := http.NewServeMux()  // Returns Router

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	// setup server
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	var PORT = strings.Split(cfg.Addr, ":")[1];

	slog.Info("Server Started on", slog.String("PORT", PORT))
	// fmt.Printf("Server Started on %s", PORT)
	
	done := make(chan os.Signal, 1)

	// For Gracefull Shutdown
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe() // Blocking
	
		if(err != nil) {
			log.Fatal("Failed to start server")
		}
	}()

	<- done

	// Stop Server
	slog.Info("Shutting down the server");


	ctx , cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	defer cancel() // As the function ends this will be called


	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

	// server.Shutdown()  => This will Shutdown your server Gracefully but it might happent that it may wait to finish the exisisting process due to which it may come into infinite loop

}
