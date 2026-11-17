package http

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var previewTripRequest previewTripRequest

	if err := json.NewDecoder(r.Body).Decode(&previewTripRequest); err != nil {
		http.Error(w, "error parsing trip preview request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	route, err := s.Service.GetRoute(ctx, &previewTripRequest.Pickup, &previewTripRequest.Destination)
	if err != nil {
		http.Error(w, "cannot fetch route from OSRM", http.StatusInternalServerError)
	}
	writeJson(w, http.StatusAccepted, route)
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
