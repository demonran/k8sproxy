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
	mu.Lock()
	defer mu.Unlock()

	response := make([]types.ClientInfo, 0, len(clients))
	for _, info := range clients {
		if info != nil {
			response = append(response, *info)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("JSON encoding failure: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "System error",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
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

func Unregister(w http.ResponseWriter, r *http.Request) {
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)

	mu.Lock()
	if _, exists := clients[clientIP]; exists {
		delete(clients, clientIP)
		log.Printf("client side%s logged out", clientIP)
	} else {
		log.Printf("Logout request failed: client side%s does not exist", clientIP)
	}
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
