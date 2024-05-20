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
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+os.Args[1])
	if err != nil {
		log.Fatalf("[%s] net.ResolveUDPAddr, err: %v", "client", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("[%s] net.DialUDP, err: %v", "client", err)
	}
	log.Printf("Connected to server: %s\n", addr.String())

	go func() {
		defer conn.Close()

		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter text: ")
			text, _ := reader.ReadString('\n')

			_, err = conn.Write([]byte(text))
			if err != nil {
				log.Fatalf("[%s] net.Write, err: %v", "client", err)
			}
		}
	}()

	go func() {
		defer conn.Close()

		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatalf("[%s] bufio.NewReader, err: %v", "client", err)
			}

			log.Printf("Received from server: %s", data)
		}
	}()

	runtime.Goexit()
}
