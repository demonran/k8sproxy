package conn

import (
	"fmt"
	"github.com/wzshiming/socks5"
	"github.com/wzshiming/sshproxy"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
	"k8sproxy/internal/route"
	"k8sproxy/pkg/options"
	"k8sproxy/pkg/util"
	"k8sproxy/tun"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	LocalHost  = "127.0.0.1"
	Socks5Port = 2223
)

func Connect() error {
	defer Unregister()
	if err := checkPermission(); err != nil {
		return err
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	opts := options.GetOption()
	socks5Address := fmt.Sprintf("%s:%d", LocalHost, Socks5Port)
	if err := startSocks5Connection(socks5Address); err != nil {
		return err
	}

	if err := tun.Ins().CheckContext(); err != nil {
		return err
	}

	if err := tun.Ins().ToSocks(socks5Address); err != nil {
		log.Printf("Failed to set up SOCKS connection: %v", err)
		return err
	}
	if err := route.SetRoute(tun.Ins(), opts.Routes); err != nil {
		log.Printf("Failed to set route: %v", err)
		return err
	}
	s := <-ch
	log.Printf("signal is %s", s)
	return fmt.Errorf("program termination")
}

func startSocks5Connection(socks5Address string) error {
	log.Println("start socks5 connection")
	var res = make(chan error)
	var ticker *time.Ticker
	gone := false
	opt := options.GetOption()
	sshHost := fmt.Sprintf("%s:%d", opt.Proxy.SshHost, 22)
	sshUser := opt.Proxy.SshUser
	sshPwd := opt.Proxy.SshPwd

	go func() {
		err := startSocksProxy(sshHost, sshUser, sshPwd, socks5Address)
		if !gone {
			res <- err
		}
		if ticker != nil {
			ticker.Stop()
		}
		time.Sleep(10 * time.Second)
		_ = startSocks5Connection(socks5Address)

	}()
	select {
	case err := <-res:
		return err
	case <-time.After(1 * time.Second):
		ticker = setupSocks5HeartBeat(sshHost, socks5Address)
		gone = true
		return nil
	}

}

func startSocksProxy(sshHost, sshUser, sshPwd, socks5Address string) error {
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	dialer, err := sshproxy.NewDialerWithConfig(sshHost, sshConfig)
	if err != nil {
		return err
	}

	defer dialer.Close()

	server := &socks5.Server{
		Logger:    log.Default(),
		ProxyDial: dialer.DialContext,
	}

	return server.ListenAndServe("tcp", socks5Address)

}

func setupSocks5HeartBeat(sshAddress, socks5Address string) *time.Ticker {
	dialer, err := proxy.SOCKS5("tcp", socks5Address, nil, proxy.Direct)
	if err != nil {
		log.Fatal("error")
	}
	ticker := time.NewTicker(60 * time.Second)
	go func() {
	TickLoop:
		for {
			select {
			case <-ticker.C:
				if c, err2 := dialer.Dial("tcp", sshAddress); err2 != nil {
					log.Printf("Heart beat err: %s", err2.Error())
				} else {
					_ = c.Close()
				}
			case <-time.After(2 * 60 * time.Second):
				log.Printf("Socks proxy heartbeat stopped")
				break TickLoop
			}

		}
	}()
	return ticker
}

func checkPermission() error {
	if !util.IsRunAsAdmin() {
		if !util.IsRunAsAdmin() {
			if util.IsWindows() {
				return fmt.Errorf("permission declined, please re-run as Administrator")
			}
			return fmt.Errorf("permission declined, please re-run with 'sudo'")
		}
	}
	return nil
}

func Unregister() {
	if err := options.UnregisterClient(); err != nil {
		log.Printf("An error occurred during the logout process: %v", err)
	}
}
