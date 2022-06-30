package main

import (
	"context"
	"fmt"
	"imageupload/config"
	"imageupload/pkg/utils/files"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const (

	// Timeouts to prevent slow-loris attacks and for avoiding build-ups
	serverReadTimeout  = 10 * time.Second
	serverIdleTimeout  = 10 * time.Second
	serverWriteTimeout = 10 * time.Second
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	files.CreateDir(config.LocalMemPath)
	files.CreateDir(config.LocalImagesDirectory)

	r := gin.New()
	r.Static("/home", "./public")
	http.Handle("/", r)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	// Set up the http server
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", config.HttpServerPort),
		ReadTimeout:  serverReadTimeout,
		IdleTimeout:  serverIdleTimeout,
		WriteTimeout: serverWriteTimeout,
	}

	go srv.ListenAndServe()

	// Block for exit
	<-termChan

	// Cleanup
	srv.Shutdown(ctx)
	cancelFunc()
}
