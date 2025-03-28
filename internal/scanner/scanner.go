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
		fmt.Printf("Port %d is closed (%s)\n", port, GetServiceByPort(port))
		return
	}

	conn.Close()

	fmt.Printf("Port %d is open (%s)\n", port, GetServiceByPort(port))
}

func ScanUDP(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	
	conn, err := net.DialTimeout("udp", address, 2*time.Second)
	if err != nil {
		fmt.Printf("Port %d is closed (%s)\n", port, GetServiceByPort(port))
		return
	}
	defer conn.Close()

	// j'envoie un paquet vide pour tester la connexion
	_, _ = conn.Write([]byte(""))

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)

	service := GetServiceByPort(port)
	
	switch {
	case err == nil:
		fmt.Printf("Port %d is open (%s)\n", port, service)
		
	case isTimeout(err):
		fmt.Printf("Port %d is open|filtred (%s)\n", port, service)
		
	default:
		fmt.Printf("Port %d is closed (%s)\n", port, service)
	}
}

func isTimeout(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}
func GetServiceByPort(port int) string {
	if service, exists := PortMapping[port]; exists {
		return service
	}
	return "Unknown"
}