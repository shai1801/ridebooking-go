package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	//TODO: Call trip service

	response := contracts.APIResponse{Data: "Ok"}

	if err := writeJSON(w, http.StatusAccepted, response); err != nil {
		http.Error(w, "failed to parse the response:", http.StatusInternalServerError)
		return
	}

}
