package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// InterfaceStats stores the statistics for an interface
type InterfaceStats struct {
	Interface string  `json:"interface"`
	Sent      float64 `json:"sent"`
	Received  float64 `json:"received"`
}

// IPStats stores the statistics for an IP address
type IPStats struct {
	IP       string  `json:"ip"`
	Sent     float64 `json:"sent"`
	Received float64 `json:"received"`
}

func main() {
	http.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request) {
		interfaceStats, ipStats := readNetworkStats()
		stats := struct {
			Interfaces []InterfaceStats `json:"interfaces"`
			IPs        []IPStats        `json:"ips"`
		}{
			Interfaces: interfaceStats,
			IPs:        ipStats,
		}
		json.NewEncoder(w).Encode(stats)
	})

	http.ListenAndServe(":8080", nil)
}

// readNetworkStats reads network interface statistics
// readNetworkStats reads network interface statistics

func readNetworkStats() ([]InterfaceStats, []IPStats) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	defer file.Close()

	var interfaceStats []InterfaceStats
	var ipStats []IPStats

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, ":") {
			continue
		}

		interfaceName, txBytes, rxBytes := parseNetworkStats(line)
		if interfaceName == "" {
			continue
		}
		if interfaceName == "lo" {
			continue
		}

		if txBytes == 0 && rxBytes == 0 {
			continue
		}

		updateInterfaceStats(&interfaceStats, interfaceName, txBytes, rxBytes)
		updateIPStats(&ipStats, interfaceName, txBytes, rxBytes)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	return interfaceStats, ipStats
}

// parseNetworkStats parses network interface statistics from a line
func parseNetworkStats(line string) (string, int, int) {
	fields := strings.Fields(line)
	if len(fields) < 17 {
		return "", 0, 0
	}

	interfaceName := strings.TrimSuffix(fields[0], ":")
	rxBytes, _ := strconv.Atoi(fields[1])
	txBytes, _ := strconv.Atoi(fields[9])

	return interfaceName, rxBytes, txBytes
}

// updateInterfaceStats updates the interface statistics
func updateInterfaceStats(interfaceStats *[]InterfaceStats, name string, txBytes, rxBytes int) {
	stats := InterfaceStats{
		Interface: name,
		Sent:      float64(txBytes),
		Received:  float64(rxBytes),
	}
	*interfaceStats = append(*interfaceStats, stats)
}

// updateIPStats updates the IP statistics
func updateIPStats(ipStats *[]IPStats, interfaceName string, txBytes, rxBytes int) {
	addrs, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return
	}
	interfaceAddrs, _ := addrs.Addrs()
	for _, addr := range interfaceAddrs {
		ip := addr.(*net.IPNet).IP.String()
		stats := IPStats{
			IP:       ip,
			Sent:     float64(txBytes),
			Received: float64(rxBytes),
		}
		*ipStats = append(*ipStats, stats)
	}
}
