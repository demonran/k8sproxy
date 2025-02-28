package types

// ProxyInfo represents the proxy information
type ProxyInfo struct {
	SshHost string `json:"sshHost"`
	SshUser string `json:"sshUser"`
	SshPwd  string `json:"sshPwd"`
}

// Option represents the proxy and route information
type Option struct {
	Proxy  ProxyInfo `json:"proxy"`
	Routes []string  `json:"routes"`
}

// ClientInfo represents the client information
type ClientInfo struct {
	ClientIP     string `json:"clientIP"`
	ClientUser   string `json:"clientUser"`
	ClientSystem string `json:"clientSystem"`
}