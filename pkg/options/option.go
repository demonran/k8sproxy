package options

type Option struct {
	PodCidr     string
	SvcCidr     string
	SshAddr     string
	SshUser     string
	SshPassword string
	ProxyPort   int
	ProxyAddr   string
}

var opt *Option

func GetOption() *Option {
	if opt == nil {
		opt = &Option{}
	}
	return opt
}

func OptionFlags() []OptionConfig {
	flags := []OptionConfig{

		{
			Target:       "PodCidr",
			DefaultValue: "10.233.64.0/18",
			Description:  "Disable access to pod IP address",
		},

		{
			Target:       "SvcCidr",
			DefaultValue: "10.233.0.0/18",
			Description:  "Specify extra IP ranges which should be route to cluster, e.g. '172.2.0.0/16', use ',' separated",
		},
		{
			Target:       "SshAddr",
			DefaultValue: "172.30.3.56",
			Description:  "ssh address",
		},
		{
			Target:       "SshUser",
			DefaultValue: "user",
			Description:  "ssh user",
		},
		{
			Target:       "SshPassword",
			DefaultValue: "password",
			Description:  "ssh password",
		},
		{
			Target:       "ProxyPort",
			DefaultValue: 2223,
			Description:  "(tun2socks mode only) Specify the local port which socks5 proxy should use",
		},
		{
			Target:       "ProxyAddr",
			DefaultValue: "127.0.0.1",
			Description:  "(tun2socks mode only) Specify the ip address or hostname which socks5 proxy should use",
		},
	}
	return flags
}
