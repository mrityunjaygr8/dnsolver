package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mrityunjaygr8/dnsolver/query"
)

func main() {
	query := query.New("www.example.com", query.TYPE_A)

	fmt.Printf("% x\n", query)

	udpServer, err := net.ResolveUDPAddr("udp", "8.8.8.8:53")
	if err != nil {
		log.Fatal("err occurred: ", err)
	}

	dialer, err := net.DialUDP("udp", nil, udpServer)
	if err != nil {
		log.Fatal("err occurred: ", err)
	}

	defer dialer.Close()

	_, err = dialer.Write(query)
	if err != nil {
		log.Fatal("err occurred: ", err)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = dialer.Read(received)
	if err != nil {
		log.Fatal("Read data failed:", err.Error())
	}

	log.Printf("% x\n", received)
}
