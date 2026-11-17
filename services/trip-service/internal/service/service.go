package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TripRepository
}

func NewService(repo domain.TripRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateTrip(ctx context.Context, fare *domain.RideFareModel) (*domain.TripModel, error) {

	t := &domain.TripModel{
		ID:       primitive.NewObjectID(),
		UserID:   fare.UserID,
		Status:   "pending",
		RideFare: fare,
	}
	return s.repo.CreateTrip(ctx, t)

}

func (s *service) GetRoute(ctx context.Context, pickup, destination *types.Coordinate) (*types.OsrmApiResponse, error) {

	responseObj := types.OsrmApiResponse{}
	//create the get osrm url with formatting using pickup and destination
	url := fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson",
		pickup.Longitude, pickup.Latitude, destination.Longitude, destination.Latitude)

	// store the reponse in variable

	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("error calling the OSRM api : %w", err)
		return nil, err
	}

	defer resp.Body.Close()

	//parse this variable in the *types.OsrmApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseObj); err != nil {
		return nil, fmt.Errorf("error parsing the OSRM response : %w", err)
	}

	return &responseObj, nil
}
