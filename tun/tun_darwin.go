package tun

import (
	"fmt"
	"k8sproxy/pkg/util"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
)

func (c *Cli) CheckContext() error {
	if !util.CanRun(exec.Command("which", "ifconfig")) {
		return fmt.Errorf("failed to found 'ifconfig' command")
	}
	if !util.CanRun(exec.Command("which", "route")) {
		return fmt.Errorf("failed to found 'route' command")
	}
	if !util.CanRun(exec.Command("which", "netstat")) {
		return fmt.Errorf("failed to found 'netstat' command")
	}
	return nil
}

var tunName string

func (c *Cli) SetRoute(ipRange []string) error {
	var lastErr error
	for i, cidr := range ipRange {
		log.Printf("Add route to %s", cidr)
		tunIp := strings.Split(cidr, "/")[0]
		if i == 0 {
			_, _, err := util.RunAndWait(exec.Command("ifconfig", c.GetName(), "inet", cidr, tunIp))
			lastErr = err
		} else {
			_, _, err := util.RunAndWait(exec.Command("ifconfig", c.GetName(), "add", cidr, tunIp))
			lastErr = err
		}
		if lastErr != nil {
			log.Printf("Failed to add ip addr %s to tun device", tunIp)
			continue
		}
		_, _, err := util.RunAndWait(exec.Command("route", "add", "-net", cidr, "-interface", c.GetName()))
		lastErr = err
		if lastErr != nil {
			log.Printf("Failed to set route %s to tun device", cidr)
		}
	}
	return lastErr
}

func (c *Cli) GetName() string {
	if tunName != "" {
		return tunName
	}
	if interfaces, err := net.Interfaces(); err == nil {
		tunNum := 0
		for _, i := range interfaces {
			if strings.HasPrefix(i.Name, util.TunNameMac) {
				if num, err := strconv.Atoi(strings.TrimPrefix(i.Name, util.TunNameMac)); err == nil && num > tunNum {
					tunNum = num
				}
			}
			tunName = fmt.Sprintf("%s%d", util.TunNameMac, tunNum+1)
		}
	}
	return tunName

}
