package tun

import (
	"fmt"
	"github.com/xjasonlyu/tun2socks/v2/engine"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (c *Cli) ToSocks(sockAddr string) error {
	tunSignal := make(chan error)

	go func() {
		var key = new(engine.Key)
		key.Proxy = sockAddr
		key.Device = fmt.Sprintf("tun://%s", c.GetName())
		key.LogLevel = "debug"
		engine.Insert(key)
		tunSignal <- engine.Start()

		defer func() {
			_ = engine.Stop()
			log.Printf("Stop tun device %s stopped", key.Device)
		}()
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
	}()

	return <-tunSignal
}
