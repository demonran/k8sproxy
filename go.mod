module k8sproxy

go 1.19

require (
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/wzshiming/socks5 v0.4.2
	github.com/wzshiming/sshproxy v0.4.3
	github.com/xjasonlyu/tun2socks/v2 v2.4.1
	golang.org/x/crypto v0.3.0
	golang.org/x/net v0.2.0
	golang.org/x/sys v0.2.0
	golang.zx2c4.com/wintun v0.0.0-20211104114900-415007cec224
)

require (
	github.com/Dreamacro/go-shadowsocks2 v0.1.7 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.0 // indirect
	github.com/go-chi/render v1.0.1 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/wzshiming/sshd v0.2.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	golang.zx2c4.com/wireguard v0.0.0-20220318042302-193cf8d6a5d6 // indirect
	gvisor.dev/gvisor v0.0.0-20220405222207-795f4f0139bb // indirect
)

replace github.com/xjasonlyu/tun2socks/v2 v2.4.1 => github.com/linfan/tun2socks/v2 v2.4.2-0.20220501081747-6f4a45525a7c
