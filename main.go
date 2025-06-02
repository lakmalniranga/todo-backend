package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lakmalniranga/todo-backend/configs"
	"github.com/lakmalniranga/todo-backend/routes"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// It's okay if .env doesn't exist, we'll use defaults or env vars
		log.Println("No .env file found, using environment variables or defaults")
	}

	// Connect to MongoDB
	configs.ConnectDB()
	defer configs.DisconnectDB()

	// Create a new router
	router := mux.NewRouter()

	// Register routes
	routes.TodoRoutes(router)

	// Add middleware for logging
	router.Use(loggingMiddleware)

	// Add health check endpoints for Kubernetes probes
	// Liveness probe - checks if the server is running
	router.HandleFunc("/health/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "API is alive")
	}).Methods("GET")

	// Readiness probe - checks if the application is ready to receive traffic
	router.HandleFunc("/health/ready", func(w http.ResponseWriter, r *http.Request) {
		// Check MongoDB connection
		err := configs.Client.Ping(context.Background(), nil)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Database connection failed")
			return
		}
		
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "API is ready to receive traffic")
	}).Methods("GET")

	// Legacy health check endpoint for backward compatibility
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "API is running")
	}).Methods("GET")

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Set up the server
	serverAddr := ":" + port
	server := &http.Server{
		Addr:         serverAddr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a goroutine
	go func() {
		fmt.Printf("Server is running on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port %s: %v\n", port, err)
		}
	}()

	// Set up graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Shutdown the server
	server.Shutdown(ctx)
	fmt.Println("Server gracefully stopped")
}

// loggingMiddleware logs all requests with their path and the time it took to process
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Printf("[%s] %s %s - Completed in %v\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, time.Since(start))
	})
}