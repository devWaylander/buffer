package main

import (
	"buffer/config"
	"buffer/handler"
	"buffer/middleware"
	"buffer/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Считываем конфиг
	config := config.Configure()

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	// Service
	service := service.New()

	// Handler
	mux := http.NewServeMux()
	handler.Configure(ctx, mux, service)
	authWrap := middleware.NewAuth(mux)
	wrappedMux := middleware.NewRequestLogger(authWrap)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: wrappedMux,
	}

	if config.GenMock == "TRUE" {
		go service.MockSaveFact(ctx)
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.Printf("Server is up on PORT: %s", config.Port)
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("Server is shutting down: %s", config.Port)
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \\n", err)
	}
}
