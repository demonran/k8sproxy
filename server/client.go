package server

import (
	"encoding/json"
	"k8sproxy/pkg/types"
	"log"
	"net"
	"net/http"
	"sync"
)

var (
	clients map[string]*types.ClientInfo
	mu      sync.Mutex
)

func init() {
	clients = make(map[string]*types.ClientInfo)
}

func GetClients(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(clients)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("register client...")
	var clientInfo types.ClientInfo
	err := json.NewDecoder(r.Body).Decode(&clientInfo)
	if err != nil {
		log.Printf("Failed to decode client info: %v", err)
		http.Error(w, "Failed to decode client info", http.StatusBadRequest)
		return
	}

	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		clientIP = r.RemoteAddr
	}

	mu.Lock()
	clients[clientIP] = &clientInfo
	mu.Unlock()

	opt := *GetConfig()

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(opt)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
