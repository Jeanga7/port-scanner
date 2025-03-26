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

// ScanUDP vérifie si un port est ouvert en UDP
func ScanUDP(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("udp", address, 2*time.Second)

	if err != nil {
		fmt.Printf("Port %d is closed (%s)\n", port, GetServiceByPort(port))
		return
	}
	defer conn.Close()

	// pour verification supp, j'envoie un paquet UDP vide
	_, err = conn.Write([]byte(""))
	if err != nil {
		fmt.Printf("Port %d is closed (%s)\n", port, GetServiceByPort(port))
		return
	}

	//definition d'un delai pour la reponse
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		//si pas de reponse, on suppose que le port est fermé
		fmt.Printf("Port %d is closed (%s)\n", port, GetServiceByPort(port))
		return
	}

	fmt.Printf("Port %d is open (%s)\n", port, GetServiceByPort(port))
}

func GetServiceByPort(port int) string {
	if service, exists := PortMapping[port]; exists {
		return service
	}
	return "Unknown"
}