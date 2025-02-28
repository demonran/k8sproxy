package main

import (
	"encoding/json"
	"k8sproxy/pkg/types"
	"k8sproxy/server"
	"log"
	"net"
	"net/http"
	"sync"
)

// 定义一个全局变量来存储客户端信息
var mu sync.Mutex

func init() {
	server.Clients = make(map[string]*types.ClientInfo)
}

func getOptions(w http.ResponseWriter, r *http.Request) {
	log.Printf("register client...")
	var clientInfo types.ClientInfo
	err := json.NewDecoder(r.Body).Decode(&clientInfo)
	if err != nil {
		log.Printf("Failed to decode client info: %v", err)
		http.Error(w, "Failed to decode client info", http.StatusBadRequest)
		return
	}

	// 提取客户端的IP地址，去掉端口
	clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Failed to split host and port: %v", err)
		clientIP = r.RemoteAddr // 如果解析失败，使用原始值
	}
	clientInfo.ClientIP = clientIP
	clientUser := clientInfo.ClientUser
	clientSystem := clientInfo.ClientSystem

	log.Printf("Received client info - IP: %s, User: %s, System: %s", clientIP, clientUser, clientSystem)

	// 存储客户端信息
	mu.Lock()
	server.Clients[clientIP] = &clientInfo
	mu.Unlock()

	// 根据客户端信息返回代理信息
	opt := types.Option{
		Proxy: types.ProxyInfo{
			SshHost: "172.30.3.50",
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

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(opt)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/options", getOptions)
	http.HandleFunc("/clients", server.GetClients)
	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
