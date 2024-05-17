package main

import (
	"log"
	"net"
	"runtime"
	"time"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP address: %v", err)
	}

	go func() {
		defer conn.Close()

		for {
			var buf [1024]byte
			n, addr, err := conn.ReadFromUDP(buf[:])
			if err != nil {
				log.Printf("Failed to read from UDP address: %v", err)
				continue
			}

			log.Printf("Received %d bytes from %v: %s", n, addr, string(buf[:n]))

			time.Sleep(3 * time.Second)
			conn.WriteToUDP([]byte(string(buf[:n])+"\n"), addr)
		}
	}()

	runtime.Goexit()

	// for {
	// 	var buf [1024]byte
	// 	n, addr, err := conn.ReadFromUDP(buf[:])
	// 	if err != nil {
	// 		log.Printf("Failed to read from UDP address: %v", err)
	// 		continue
	// 	}
	// 	log.Printf("Received %d bytes from %v: %s", n, addr, string(buf[:n]))

	// 	conn.WriteToUDP([]byte("Hello, client!\n"), addr)

	// 	time.Sleep(5 * time.Second)

	// 	conn.WriteToUDP([]byte("A!\n"), addr)
	// }
}
