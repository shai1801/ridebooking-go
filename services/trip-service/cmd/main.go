package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {

	ctx := context.Background()

	repo := repository.NewInmemRepository()
	svc := service.NewService(repo)

	fare := &domain.RideFareModel{}
	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Print(t)

	for {
		time.Sleep(2 * time.Second)
	}
}
