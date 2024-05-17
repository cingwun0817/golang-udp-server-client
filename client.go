package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Failed to dial UDP address: %v", err)
	}

	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		log.Fatalf("Failed to write to UDP address: %v", err)
	}

	go func() {
		defer conn.Close()

		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')
			fmt.Println(text)

			_, err = conn.Write([]byte(text))
			if err != nil {
				log.Fatalf("Failed to write to UDP address: %v", err)
			}
		}
	}()

	go func() {
		defer conn.Close()

		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatalf("Failed to read from UDP address: %v", err)
			}

			log.Printf("Received from server: %s", data)
		}
	}()

	runtime.Goexit()

	// msg := make([]byte, 512)
	// for {
	// 	n, err := conn.Read(msg)
	// 	if err != nil {
	// 		log.Fatalf("Failed to read from UDP address: %v", err)
	// 	}

	// 	log.Printf("server replied with: %s \n", string(msg[:n]))
	// }

	// for {
	// 	data, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		log.Fatalf("Failed to read from UDP address: %v", err)
	// 	}
	// 	log.Printf("Received from server: %s", data)
	// }
}
