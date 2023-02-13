package tun

import (
	"fmt"
	"golang.zx2c4.com/wintun"
	"k8sproxy/pkg/util"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

func (c *Cli) CheckContext() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("failed to found tun driver: %v", r)
		}
	}()
	if !util.CanRun(exec.Command("netsh")) {
		return fmt.Errorf("failed to found 'netsh' command")
	}
	wintun.RunningVersion()
	return nil
}

func (c *Cli) GetName() string {
	return "K8sProxyTunnel"
}

func (c *Cli) SetRoute(ipRange []string) error {
	var lastErr error
	for i, cidr := range ipRange {
		log.Printf("Add route to %s", cidr)
		_, mask, err := toIpAndMask(cidr)
		tunIp := strings.Split(cidr, "/")[0]
		if i == 0 {
			_, _, err := util.RunAndWait(exec.Command("netsh",
				"interface",
				"ipv4",
				"set",
				"address",
				c.GetName(),
				"static",
				tunIp,
				mask))
			lastErr = err
		} else {
			_, _, err = util.RunAndWait(exec.Command("netsh",
				"interface",
				"ipv4",
				"add",
				"address",
				c.GetName(),
				tunIp,
				mask,
			))
			lastErr = err
		}
		if lastErr != nil {
			log.Printf("Failed to add ip addr %s to tun device", tunIp)
			continue
		}
		_, _, err = util.RunAndWait(exec.Command("netsh",
			"interface",
			"ipv4",
			"add",
			"route",
			cidr,
			c.GetName(),
			tunIp,
		))
		lastErr = err
		if lastErr != nil {
			log.Printf("Failed to set route %s to tun device", cidr)
		}
	}

	return lastErr
}

func toIpAndMask(cidr string) (string, string, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", "", err
	}
	val := make([]byte, len(ipNet.Mask))
	copy(val, ipNet.Mask)

	var s []string
	for _, i := range val[:] {
		s = append(s, strconv.Itoa(int(i)))
	}
	return ipNet.IP.String(), strings.Join(s, "."), nil
}
