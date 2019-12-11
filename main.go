package main

import (
	"LedImageView-WebAPI/pkg/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := router.InitRouter()
	server := http.Server{
		Addr:              "0.0.0.0:5000",
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	_ = server.ListenAndServe()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<- quit
	_ = server.Close()

}