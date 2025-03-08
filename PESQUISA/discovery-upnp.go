package main

import (
	"fmt"
	"net"
	"time"
)

const ssdpDiscover = `M-SEARCH * HTTP/1.1
HOST: 239.255.255.250:1900
MAN: "ssdp:discover"
MX: 3
ST: ssdp:all
`

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.WriteTo([]byte(ssdpDiscover), addr)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 2048)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	for {
		n, _, err := conn.ReadFrom(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
