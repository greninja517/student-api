package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/greninja517/student-api/internal/config"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	response := []byte("Welcome to Student API Home Page")
	_, err := w.Write(response)
	if err != nil {
		log.Println("Failed to write the response. Retrying...")
	}
}

func main() {
	// loading the server configuration
	cfg := config.ConfigurationLoader()

	// setting up the router
	router := http.NewServeMux()
	router.HandleFunc("/", HomePage)

	// setting up the server
	server := &http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	go func() {
		log.Println("Server running at ", cfg.Address)
		err := server.ListenAndServe()
		if err != nil {
			slog.Info("", slog.String("Error", err.Error()))
		}
	}()

	// graceful shutdown of server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop // wait for the interrupt signal

	fmt.Println("Shutting Down the Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		// slog is used for displayed the logs with key-value pair as well
		slog.Error("Failed to Shutdown", slog.String("Error: ", err.Error()))
	}
	slog.Info("Server shut down success...")
}
