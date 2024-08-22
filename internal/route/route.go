package route

import "k8sproxy/tun"

func SetRoute(tun tun.Tunnel, ipRange []string) error {
	return tun.SetRoute(ipRange)
}
