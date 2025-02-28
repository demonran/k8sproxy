package server

import (
	"encoding/json"
	"k8sproxy/pkg/types"
	"log"
	"net/http"
)

var Clients map[string]*types.ClientInfo

func GetClients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Clients)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
