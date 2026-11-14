package main

import (
	"fmt"
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")
	mux := http.NewServeMux()
	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	s := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	if err := s.ListenAndServe(); err != nil {
		err = fmt.Errorf("error in starting server: %w", err)
		log.Print(err)
	}
}
