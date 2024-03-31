package api

import (
	"encoding/json"
	"log"
	"net/http"
	"search-server/types"
	"strings"
)

func ImageDataHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	appData, ok := ctx.Value(types.AppDataKey{}).(types.JsonData)

	if !ok {
		http.Error(w, "Failed to retrieve appData", http.StatusInternalServerError)
		return
	}
	// Extract ID from the request URL
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	// Retrieve ImageData for the given ID
	imageData, found := appData.DocumentInfoMap[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	// Marshal ImageData to JSON
	responseJSON, err := json.Marshal(imageData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error marshaling JSON:", err)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write response
	w.Write(responseJSON)
}
