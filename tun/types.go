package tun

type Tunnel interface {
	CheckContext() error
	ToSocks(sockAddr string) error
	SetRoute(ipRange []string) error
	GetName() string
}

type Cli struct{}

var instance *Cli

func Ins() Tunnel {
	if instance == nil {
		instance = &Cli{}
	}
	return instance
}
