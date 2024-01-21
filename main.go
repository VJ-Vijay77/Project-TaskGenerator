package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"taskgenerator/processingnode"
	"taskgenerator/taskgenerator"
)

func main() {
	// Create a context for signaling shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize WaitGroup for graceful shutdown
	var wg sync.WaitGroup

	// Start Task Generator
	go taskgenerator.RunGenerateTasks(ctx)
	wg.Add(1)

	// Start Processing Node
	processingnode.RunProcessingNode(&wg, ctx)

	// Handle graceful shutdown with a signal
	handleShutdownGracefully(cancel)

	// Wait for all Goroutines to finish before exiting
	wg.Wait()
}

func handleShutdownGracefully(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("Received shutdown signal. Initiating graceful shutdown...")

		// Cancel the context to signal termination to all Goroutines
		cancel()

		fmt.Println("Shutdown complete. Exiting.")
		os.Exit(0)
	}()
}
