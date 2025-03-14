package server

import (
	"encoding/json"
	"k8sproxy/pkg/types"
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
