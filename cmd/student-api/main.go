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
	"github.com/greninja517/student-api/internal/http/handlers/student"
	"github.com/greninja517/student-api/internal/storage/sqlite"
)

func main() {
	// loading the server configuration
	cfg := config.ConfigurationLoader()

	//setting up the Database
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("Failed to connect to Database. Error: ", err)
	}
	slog.Info("Database Initialized.... ", slog.String("Status", "Success"))

	// setting up the router
	router := http.NewServeMux()
	router.HandleFunc("POST /students", student.Student(storage))

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
		slog.Error("Forced to Shutdown the Server...", slog.String("Error: ", err.Error()))
	}
	slog.Info("Server shut down success...")
}
