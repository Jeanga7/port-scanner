package test

import (
    "bytes"
    "io"
    "os"
    "testing"

    "github.com/Jeanga7/port-scanner/internal/scanner"
)

// TestScanTCP teste la fonction ScanTCP
func TestScanTCP(t *testing.T) {
    tests := []struct {
        host     string
        port     int
        expected string
    }{
        {"127.0.0.1", 80, "Port 80 is open (HTTP)\n"},
        {"127.0.0.1", 22, "Port 22 is open (SSH)\n"},
        {"127.0.0.1", 9999, "Port 9999 is closed (Unknown)\n"},
    }

    for _, tt := range tests {
        t.Run("TCP Scan", func(t *testing.T) {
            // Capture la sortie
            output := captureOutput(func() {
                scanner.ScanTCP(tt.host, tt.port)
            })
            if output != tt.expected {
                t.Errorf("expected %v, got %v", tt.expected, output)
            }
        })
    }
}

// TestScanUDP teste la fonction ScanUDP
func TestScanUDP(t *testing.T) {
    tests := []struct {
        host     string
        port     int
        expected string
    }{
        {"127.0.0.1", 53, "Port 53 is open (DNS)\n"},
        {"127.0.0.1", 1604, "Port 1604 is open (Unknown)\n"},
    }

    for _, tt := range tests {
        t.Run("UDP Scan", func(t *testing.T) {
            output := captureOutput(func() {
                scanner.ScanUDP(tt.host, tt.port)
            })
            if output != tt.expected {
                t.Errorf("expected %v, got %v", tt.expected, output)
            }
        })
    }
}

// TestGetServiceByPort teste la fonction GetServiceByPort
func TestGetServiceByPort(t *testing.T) {
    tests := []struct {
        port     int
        expected string
    }{
        {80, "HTTP"},
        {22, "SSH"},
        {53, "DNS"},
        {9999, "Unknown"},
    }

    for _, tt := range tests {
        t.Run("Service Detection", func(t *testing.T) {
            service := scanner.GetServiceByPort(tt.port)
            if service != tt.expected {
                t.Errorf("expected %v, got %v", tt.expected, service)
            }
        })
    }
}

// captureOutput capture la sortie d'une fonction
func captureOutput(f func()) string {
    // Crée un buffer temporaire pour capturer la sortie
    r, w, _ := os.Pipe()
    stdout := os.Stdout
    os.Stdout = w

    // Exécute la fonction
    f()

    // Restaure la sortie standard et lit le buffer
    w.Close()
    os.Stdout = stdout
    var buf bytes.Buffer
    io.Copy(&buf, r)
    return buf.String()
}