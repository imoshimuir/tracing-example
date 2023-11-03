package main

import (
	"context"
	"fmt"
	"latest-news/telemetry"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const serviceName = "news-service"
func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initalise the tracer
	tracer, err := telemetry.InitTracer(ctx, cfg.Telemetry, serviceName)
	if err != nil {
		log.Fatalf("Failed to initialize telemetry: %v", err)
	}

	echo, err := InstantiateServer(ctx, cfg, tracer)
	if err != nil {
		log.Fatalf("Failed to instantiate server: %v", err)
	}

	go func() {
		if err := echo.Start(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigs

	log.Println("Received termination signal - Shutting down server")

	if err := echo.Shutdown(ctx); err != nil {
		log.Printf("Failed to shut down server: %v", err)
	}
}