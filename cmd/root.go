package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Jeanga7/port-scanner/internal/scanner"
	"github.com/spf13/cobra"
)
var (
    tcpScan bool
    udpScan bool
    portRange string
)

// commande principale
var rootCmd = &cobra.Command{
	Use:   "active [OPTIONS] [HOST] [PORT]",
	Short: "A simple port scanner",
	Long:  `active is a tool to scan open and closed ports on a given host.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]

		if portRange == "" {
			fmt.Println("Error: You must provide a port or range of ports using -p.")
			return
		}

		ports := parsePorts(portRange)

		if len(ports) == 0 {
			fmt.Println("Error: Invalid port range.")
			return
		}

		for _, port := range ports {
			if tcpScan {
				scanner.ScanTCP(host, port)
			}
			if udpScan {
				scanner.ScanUDP(host, port)
			}
		}
	},
}

// fonction pour parser la plage de ports (eg: 80, 80-85)
func parsePorts(portStr string) []int {
	var ports []int

	if strings.Contains(portStr, "-") {
		parts := strings.Split(portStr, "-")
		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])

		if err1 == nil && err2 == nil && start < end {
			for i := start; i <= end; i++ {
				ports = append(ports, i)
			}
		}
	} else {
		port, err := strconv.Atoi(portStr)

		if err == nil {
			ports = append(ports, port)
		}
	}
	return ports
}

func customHelp(cmd *cobra.Command, args []string) {
	fmt.Println(`Usage: active [OPTIONS] [HOST] [PORT]
Options:
  -p               Range of ports to scan
  -u               UDP scan
  -t               TCP scan
  --help           Show this message and exit.`)
	os.Exit(0)
}

// Execute exécute la commande CLI
func Execute() {
	rootCmd.SetHelpFunc(customHelp)

	rootCmd.Flags().BoolVarP(&tcpScan, "tcp", "t", false, "TCP scan")
    rootCmd.Flags().BoolVarP(&udpScan, "udp", "u", false, "UDP scan")
    rootCmd.Flags().StringVarP(&portRange, "port", "p", "", "Port or range of ports to scan (ex: 80 or 20-25)")


	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
