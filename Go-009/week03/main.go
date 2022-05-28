package main

import (
	"log"
	"net/http"
	"time"

	"G20220607100065/Go-009/week03/service"
)

func main() {
	s0 := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	s1 := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: nil,
	}

	s2 := &http.Server{
		Addr:    "127.0.0.1:8082",
		Handler: nil,
	}

	logger := log.Default()

	app := service.NewApp(
		service.WithServer(s0),
		service.WithServer(s1),
		service.WithServer(s2),
		service.WithLog(logger),
	)

	time.AfterFunc(time.Second*3, func() {
		err := app.Stop()
		if err != nil {
			return
		}
	})

	if err := app.Run(); err != nil {
		log.Printf("%v", err)
	}

	logger.Println("app service has exited")
}
