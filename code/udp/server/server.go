package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var addr = flag.String("addr", ":10000", "udp server bing address")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	//Resolving address
	udpAddr, err := net.ResolveUDPAddr("udp", *addr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Build listining connections
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer conn.Close()

	// Interacting with one client at a time
	recvBuff := make([]byte, 1024)
	for {
		log.Println("Ready to receive packets!")
		// Receiving a message
		rn, rmAddr, err := conn.ReadFromUDP(recvBuff)
		if err != nil {
			log.Println("Error:", err)
			return
		}

		fmt.Printf("<<< Packet received from: %s, data: %s\n", rmAddr.String(), string(recvBuff[:rn]))
		// Sending the same message back to current client
		_, err = conn.WriteToUDP(recvBuff[:rn], rmAddr)
		if err != nil {
			log.Println("Error:", err)
			return
		}
		fmt.Println(">>> Sent packet to: ", rmAddr.String())
	}
}

//https://gocn.io/article/371
