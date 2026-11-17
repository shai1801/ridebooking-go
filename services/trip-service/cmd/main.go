package main

import (
	"log"
	"net/http"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {

	repo := repository.NewInmemRepository()
	svc := service.NewService(repo)
	mux := http.NewServeMux()
	httpHandler := h.HttpHandler{Service: svc}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := http.Server{
		Handler: mux,
		Addr:    ":8083",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("http server error : %v", err)
	}
}
