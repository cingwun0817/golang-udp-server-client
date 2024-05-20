package main

import (
	"log"
	"net"
	"os"
	"runtime"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+os.Args[1])
	if err != nil {
		log.Fatalf("[%s] net.ResolveUDPAddr, err: %v", "server", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("[%s] net.ListenUDP, err: %v", "server", err)
	}
	log.Printf("%s Listening...", "127.0.0.1:"+os.Args[1])

	go func() {
		defer conn.Close()

		addrClients := make(map[string]*net.UDPAddr)
		for {
			var buf [1024]byte
			n, addr, err := conn.ReadFromUDP(buf[:])
			if err != nil {
				log.Fatalf("[%s] conn.ReadFromUDP, err: %v", "server", err)
			}

			if _, ok := addrClients[addr.String()]; !ok {
				addrClients[addr.String()] = addr

				log.Printf("New client: %s\n", addr.String())
			}

			log.Printf("Received %s: %s", addr.String(), string(buf[:n]))

			for _, addrClient := range addrClients {
				conn.WriteToUDP([]byte(string(buf[:n])+"\n"), addrClient)
			}

			log.Printf("Broadcast: %s", string(buf[:n]))
		}
	}()

	runtime.Goexit()
}
