package main

import (
	"flag"
	"k8sproxy/server"
	"log"
	"net/http"
)

func main() {
	configPath := flag.String("config", "./config.json", "Path to config file")
	flag.Parse()

	server.LoadConfig(*configPath)

	http.HandleFunc("/options", server.Register)
	http.HandleFunc("/clients", server.GetClients)
	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
