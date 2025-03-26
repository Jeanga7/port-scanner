package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanTCP vérifie si un port est ouvert en TCP
func ScanTCP(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err != nil {
		fmt.Printf("Port %d is closed\n", port)
		return
	}

	conn.Close()
	fmt.Printf("Port %d is open\n", port)
}

// ScanUDP vérifie si un port est ouvert en UDP
func ScanUDP(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("udp", address, 2*time.Second)

	if err != nil {
		fmt.Printf("Port %d is closed\n", port)
		return
	}

	conn.Close()
	fmt.Printf("Port %d is open\n", port)
}
