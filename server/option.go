package server

import (
	"encoding/json"
	"k8sproxy/pkg/types"
	"net/http"
	"os"
)

var config *types.Option

func LoadConfig(configPath string) {
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = &types.Option{}
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *types.Option {
	return config
}

// 新增 CORS 中间件
func EnableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
