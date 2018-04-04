package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	b := make([]byte, 1)
	os.Stdin.Read(b)
	conn.Write([]byte("hello"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())
}
