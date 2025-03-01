package options

import (
	"bytes"
	"encoding/json"
	"io"
	"k8sproxy/pkg/types"
	"k8sproxy/pkg/util"
	"log"
	"net/http"
	"os"
)

var opt *types.Option

func Init() {
	if opt == nil {
		opt = fetchOptionsFromServer()
	}
}

func GetOption() *types.Option {
	return opt
}

func fetchOptionsFromServer() *types.Option {
	clientIP, err := util.GetClientIP()
	if err != nil {
		log.Printf("Failed to get client IP: %v", err)
		clientIP = ""
	}
	clientUsername, err := util.GetSystemUsername()
	if err != nil {
		log.Printf("Failed to get system username: %v", err)
		clientUsername = ""
	}
	clientSystem, err := os.Hostname()
	if err != nil {
		log.Printf("Failed to get system hostname: %v", err)
		clientSystem = "" // 如果获取失败，回退到 GOOS
	}

	clientInfo := types.ClientInfo{
		ClientIP:     clientIP,
		ClientUser:   clientUsername,
		ClientSystem: clientSystem,
	}

	jsonData, err := json.Marshal(clientInfo)
	if err != nil {
		log.Printf("Failed to marshal client info: %v", err)
		return defaultOptions()
	}

	resp, err := http.Post("http://172.30.3.50:8080/options", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to fetch options from server: %v", err)
		return defaultOptions()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return defaultOptions()
	}

	var options types.Option
	err = json.Unmarshal(body, &options)
	if err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return defaultOptions()
	}

	return &options
}

func defaultOptions() *types.Option {
	return &types.Option{
		Proxy: types.ProxyInfo{
			SshHost: "172.30.3.56",
			SshUser: "root",
			SshPwd:  "",
		},
		Routes: []string{
			"10.233.64.0/18",
			"10.233.0.0/18",
			"10.11.0.0/16",
			"10.10.0.0/16",
		},
	}
}

func OptionFlags() []OptionConfig {
	return []OptionConfig{}
}
