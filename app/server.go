package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"
)

func RunServer(server *http.Server) {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	server.Shutdown(ctx)
	log.Println("Server gracefully stopped")
	cancel()
}

func NewServer(handler http.Handler) *http.Server {
	port := viper.GetString("port")
	return &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
}
