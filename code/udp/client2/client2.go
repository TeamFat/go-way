package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	laddr = flag.String("laddr", "127.0.0.1:9000", "local server address")
	raddr = flag.String("raddr", "127.0.0.1:10000", "remote server address")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	// Resolving Address
	localAddr, err := net.ResolveUDPAddr("udp", *laddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", *raddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Build listening connections
	conn, err := net.ListenUDP("udp", localAddr)
	// Exit if some error occured
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	// write a message to server
	_, err = conn.WriteToUDP([]byte("hello"), remoteAddr)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(">>> Packet sent to: ", *raddr)
	}

	// Receive response from server
	buf := make([]byte, 1024)
	rn, remAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("<<<  %d bytes received from: %v, data: %s\n", rn, remAddr, string(buf[:rn]))
	}
}
