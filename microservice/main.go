package main

import (
	"context"
	_ "context"
	"fmt"
	"github.com/seidu626/playground/microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "portal-api", log.LstdFlags)
	employeeHandler := handlers.NewEmployee(logger)
	homeHandler := handlers.NewHome(logger)

	mux := http.NewServeMux()
	mux.Handle("/employee", employeeHandler)
	mux.Handle("/", homeHandler)

	server := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		IdleTimeout:       300 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	go func() {
		status := server.ListenAndServe()
		if status != nil {
			fmt.Printf("%s %s", status, "An error occurred")
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Printf("Handling graceful shutdown: %s", sig)
	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalln(err)
	}

}
