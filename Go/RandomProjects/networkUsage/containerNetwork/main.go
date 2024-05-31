package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// InterfaceStats stores the statistics for an interface
type InterfaceStats struct {
	Interface string
	Sent      float64
	Received  float64
}

// IPStats stores the statistics for an IP address
type IPStats struct {
	IP       net.IP
	Sent     float64
	Received float64
}

// monitorNetworkActivity continuously monitors network activity
func monitorNetworkActivity(interval time.Duration) {
	for {
		interfaceStats, ipStats := readNetworkStats()
		//printStats(readNetworkStats())
		printInterfaceStats(interfaceStats)
		printIPStats(ipStats)

		time.Sleep(interval)
	}
}

// printStats prints the aggregated interface and individual IP address statistics
func printStats(interfaceStats []InterfaceStats, ipStats []IPStats) {
	var output strings.Builder

	for _, stats := range interfaceStats {
		if stats.Sent != 0 || stats.Received != 0 {
			output.WriteString(fmt.Sprintf("%-10s %9s\n", "Interface:", stats.Interface))
			if stats.Sent != 0 {
				output.WriteString(fmt.Sprintf("%-10s %9.fB\n", "Sent:", stats.Sent))
			}
			if stats.Received != 0 {
				output.WriteString(fmt.Sprintf("%-10s %9.fB\n", "Received:", stats.Received))
			}
		}
	}

	for _, stats := range ipStats {
		if stats.Sent != 0 || stats.Received != 0 {
			output.WriteString(fmt.Sprintf("\t%-10s %9s\n", "IP:", stats.IP))
			if stats.Sent != 0 {
				output.WriteString(fmt.Sprintf("\t%-10s %9.fB\n", "Sent:", stats.Sent))
			}
			if stats.Received != 0 {
				output.WriteString(fmt.Sprintf("\t%-10s %9.fB\n", "Received:", stats.Received))
			}
		}
	}

	// Remove the trailing newline character
	result := strings.TrimSuffix(output.String(), "\n")

	// Print the result using carriage return without newline
	fmt.Printf("\r%s", result)
}

// readNetworkStats reads network interface statistics from /proc/net/dev
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
		ip := addr.(*net.IPNet).IP
		stats := IPStats{
			IP:       ip,
			Sent:     float64(txBytes),
			Received: float64(rxBytes),
		}
		*ipStats = append(*ipStats, stats)
	}
}

// printInterfaceStats prints the aggregated interface statistics
func printInterfaceStats(interfaceStats []InterfaceStats) {
	for _, stats := range interfaceStats {
		if stats.Sent != 0 || stats.Received != 0 {
			fmt.Printf("%-10s %9s\n", "Interface:", stats.Interface)
			if stats.Sent != 0 {
				fmt.Printf("%-10s %9.fB\n", "Sent:", stats.Sent)
			}
			if stats.Received != 0 {
				fmt.Printf("%-10s %9.fB\n", "Received:", stats.Received)
			}
		}
	}
	println()
}

// printIPStats prints the individual IP address statistics
func printIPStats(ipStats []IPStats) {
	for _, stats := range ipStats {
		if stats.Sent != 0 || stats.Received != 0 {
			fmt.Printf("\t%-10s %9s\n", "IP:", stats.IP)
			if stats.Sent != 0 {
				fmt.Printf("\t%-10s %9.fB\n", "Sent:", stats.Sent)
			}
			if stats.Received != 0 {
				fmt.Printf("\t%-10s %9.fB\n", "Received:", stats.Received)
			}
		}
	}
}

func main() {
	interval := 2 * time.Second
	monitorNetworkActivity(interval)
}
