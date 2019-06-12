package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/multiplio/ozymandias/api"
	"github.com/multiplio/ozymandias/auth"
	"github.com/multiplio/ozymandias/server"
)

func main() {
	user := auth.GetUser()

	address := ":5000"
	buildPath := path.Clean("ui/build")

	mux := http.NewServeMux()
	mux.Handle("/login/", auth.Handler())
	mux.Handle("/api/", api.Handler(user))
	mux.Handle("/", server.Handler(buildPath))

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	errs := make(chan error, 1)
	go func() {
		fmt.Println("Starting", address)
		errs <- srv.ListenAndServe()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		fmt.Println("Sutting down...")
		os.Exit(0)
	case err := <-errs:
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	}

	shutdown, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdown); err != nil {
		fmt.Println("Failed to shutdown server:", err.Error())
		os.Exit(1)
	}
}
