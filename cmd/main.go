package main

import (
	"context"
	"merch-store/internal/app"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	svc, err := app.NewService()
	if err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	doneChan := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		if err := svc.Start(ctx); err != nil {
			panic(err)
		}
	}()

	go func() {
		<-signalChan
		svc.Log.Info("Received termination signal. Shutting down...", map[string]interface{}{})

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer shutdownCancel()

		if err := svc.Stop(shutdownCtx); err != nil {
			panic(err)
		}

		close(doneChan)
		cancel()
	}()

	wg.Wait()
	<-doneChan
}
