package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {

	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	////TODO: Call trip service

	tripClient, err := grpc_clients.NewTripServiceClient()

	if err != nil {
		log.Printf("error in createing grpc trip client: %v", err)
		http.Error(w, "trip service unavailable", http.StatusServiceUnavailable)
		return
	}

	defer tripClient.Close()

	resp, err := tripClient.Client.PreviewTrip(r.Context(), reqBody.toProto())

	if err != nil {
		log.Printf("error in fetching grpc Trip Preview Response: %v", err)
		http.Error(w, "failed to preview trip", http.StatusBadGateway)
		return
	}

	response := contracts.APIResponse{Data: resp}

	if err := writeJSON(w, http.StatusAccepted, response); err != nil {
		log.Printf("error writing trip preview response: %v", err)
	}

}

func handleTripStart(w http.ResponseWriter, r *http.Request) {

	var startTripRequest *startTripRequest
	if err := json.NewDecoder(r.Body).Decode(&startTripRequest); err != nil {
		log.Printf("error in unmarshalling the trip start request to object: %v", err)
	}

	defer r.Body.Close()

	tripClient, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Printf("error in creating GRPC client: %v", err)
		return
	}

	defer tripClient.Close()
	resp, err := tripClient.Client.CreateTrip(r.Context(), startTripRequest.toProto())

	if err != nil {
		log.Printf("error in getting response from start trip service: %v", err)
		http.Error(w, "failed to create trip", http.StatusBadGateway)
		return
	}

	response := contracts.APIResponse{
		Data: resp,
	}
	if err := writeJSON(w, http.StatusCreated, response); err != nil {
		http.Error(w, "error in creating trip", http.StatusBadGateway)
		return
	}
}
