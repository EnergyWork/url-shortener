package main

import (
	"context"
	"os"
	"os/signal"
	"time"
	"url_shortener/backend/lib"
)

const (
	configFile = "config.yml"
)

func main() {
	l := lib.NewLogger().SetMethod("Main")

	// panic handler
	defer func() {
		if err := recover(); err != nil {
			l.Fatal(err)
		}
	}()

	// Setup and run server

	config := lib.NewConfig(configFile)
	server := lib.NewServer(config)
	server.ConfigureRouter()
	if err := server.ConnectToDB(); err != nil {
		l.Error(err)
	}
	if err := server.Run(); err != nil {
		l.Error(err)
	}

	// Shoutdowm handle

	quit := make(chan os.Signal)
	// holding here
	// waiting for a interrupt signal
	signal.Notify(quit, os.Interrupt)
	<-quit
	l.Infof("[Control-C] Get signal: shutdown server ...")
	signal.Reset(os.Interrupt)
	l.Infof("Server shutting down")
	// context: wait for 3 seconds
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)

	defer cancel()
	// call for shutdown
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Server Shutdown failed: %v", err)
	}
	l.Infof("Server exiting")
}
