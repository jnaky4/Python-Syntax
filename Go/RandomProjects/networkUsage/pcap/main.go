package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// InterfaceStats stores the statistics for an interface
type InterfaceStats struct {
	Interface string `json:"interface"`
	Sent      int    `json:"sent"`
	Received  int    `json:"received"`
}

// IPStats stores the statistics for an IP address
type IPStats struct {
	IP       string `json:"ip"`
	Name     string `json:"name"`
	Sent     int    `json:"sent"`
	Received int    `json:"received"`
}

func main() {
	// Find all available network interfaces
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Channel to handle termination signal
	stopCapture := make(chan os.Signal, 1)
	signal.Notify(stopCapture, syscall.SIGINT, syscall.SIGTERM)

	// Mutex for synchronizing access to ipStats
	var mutex sync.Mutex

	// Map to store IP address statistics
	ipStats := make(map[string]IPStats)

	// Capture packets and update IP statistics for each network interface
	for _, device := range devices {
		if device.Name != "en0" {
			continue
		}
		if len(device.Addresses) == 0 {
			continue // Skip interfaces with no addresses
		}
		// Open the network interface for packet capture
		handle, err := pcap.OpenLive(device.Name, 1600, true, pcap.BlockForever)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		defer handle.Close()

		// Capture packets and update IP statistics
		go func(device pcap.Interface) {

			for packet := range gopacket.NewPacketSource(handle, handle.LinkType()).Packets() {
				networkLayer := packet.NetworkLayer()
				if networkLayer == nil {
					continue
				}

				srcIP := networkLayer.NetworkFlow().Src().String()
				srcDNS, _ := resolveDNSName(srcIP)
				dstIP := networkLayer.NetworkFlow().Dst().String()
				dstDNS, _ := resolveDNSName(dstIP)

				// Update IPStats for source IP
				mutex.Lock()
				ipStats[srcIP] = IPStats{
					IP:       srcIP,
					Name:     srcDNS,
					Sent:     ipStats[srcIP].Sent + len(packet.Data()),
					Received: ipStats[srcIP].Received,
				}

				// Update IPStats for destination IP
				ipStats[dstIP] = IPStats{
					IP:       dstIP,
					Name:     dstDNS,
					Sent:     ipStats[dstIP].Sent,
					Received: ipStats[dstIP].Received + len(packet.Data()),
				}
				mutex.Unlock()
			}
		}(device)

		// Periodically print IP address statistics
		go func(device pcap.Interface) {
			ticker := time.NewTicker(10 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					mutex.Lock()
					fmt.Printf("IP Address Statistics for Interface %s:\n", device.Name)
					for _, stats := range ipStats {
						printIpStats(stats, "Kb")
					}
					fmt.Println()
					// Recreate ipStats map to reset statistics
					ipStats = make(map[string]IPStats)
					mutex.Unlock()
				case <-stopCapture:
					return
				}
			}
		}(device)
	}

	// Wait for termination signal
	<-stopCapture
}

//func main() {
//	// Find all available network interfaces
//	devices, err := pcap.FindAllDevs()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Channel to handle termination signal
//	stopCapture := make(chan os.Signal, 1)
//	signal.Notify(stopCapture, syscall.SIGINT, syscall.SIGTERM)
//
//	// Mutex for synchronizing access to ipStats
//	var mutex sync.Mutex
//
//	// Map to store IP address statistics
//	ipStats := make(map[string]IPStats)
//
//	var (
//		snapshot_len int32         = 1600
//		promiscuous  bool          = true
//		timeout      time.Duration = pcap.BlockForever
//		handle       *pcap.Handle
//		// Will reuse these for each packet
//		ethLayer layers.Ethernet
//		ipLayer  layers.IPv4
//		tcpLayer layers.TCP
//	)
//
//	for _, device := range devices {
//
//		handle, err = pcap.OpenLive(device.Name, snapshot_len, promiscuous, timeout)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer handle.Close()
//
//		go func(device pcap.Interface) {
//
//			packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//			for packet := range packetSource.Packets() {
//				parser := gopacket.NewDecodingLayerParser(
//					layers.LayerTypeEthernet,
//					&ethLayer,
//					&ipLayer,
//					&tcpLayer,
//				)
//				var foundLayerTypes []gopacket.LayerType
//
//				err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
//				if err != nil {
//					fmt.Println("Trouble decoding layers: ", err.Error())
//				}
//
//				//todo add this to func
//				//for packet := range gopacket.NewPacketSource(handle, handle.LinkType()).Packets() {
//				//	networkLayer := packet.NetworkLayer()
//				//	if networkLayer == nil {
//				//		continue
//				//	}
//				//
//				//	srcIP := networkLayer.NetworkFlow().Src().String()
//				//	srcDNS, _ := resolveDNSName(srcIP)
//				//	dstIP := networkLayer.NetworkFlow().Dst().String()
//				//	dstDNS, _ := resolveDNSName(dstIP)
//				//
//				//	// Update IPStats for source IP
//				//	mutex.Lock()
//				//	ipStats[srcIP] = IPStats{
//				//		IP:       srcIP,
//				//		Name:     srcDNS,
//				//		Sent:     ipStats[srcIP].Sent + len(packet.Data()),
//				//		Received: ipStats[srcIP].Received,
//				//	}
//				//
//				//	// Update IPStats for destination IP
//				//	ipStats[dstIP] = IPStats{
//				//		IP:       dstIP,
//				//		Name:     dstDNS,
//				//		Sent:     ipStats[dstIP].Sent,
//				//		Received: ipStats[dstIP].Received + len(packet.Data()),
//				//	}
//				//	mutex.Unlock()
//				//}
//
//				//todo update with old code
//				//for _, layerType := range foundLayerTypes {
//				//	if layerType == layers.LayerTypeIPv4 {
//				//		fmt.Println("IPv4: ", ipLayer.SrcIP, "->", ipLayer.DstIP)
//				//	}
//				//
//				//	//netowrking layer
//				//	if layerType == layers.LayerTypeTCP {
//				//		fmt.Println("TCP Port: ", tcpLayer.SrcPort, "->", tcpLayer.DstPort)
//				//		fmt.Println("TCP SYN:", tcpLayer.SYN, " | ACK:", tcpLayer.ACK)
//				//	}
//				//}
//
//			}
//		}(device)
//
//		// Periodically print IP address statistics
//		go func(device pcap.Interface) {
//			ticker := time.NewTicker(10 * time.Second)
//			defer ticker.Stop()
//
//			for {
//				select {
//				case <-ticker.C:
//					mutex.Lock()
//					fmt.Printf("IP Address Statistics for Interface %s:\n", device.Name)
//					for _, stats := range ipStats {
//						printIpStats(stats, "Kb")
//					}
//					fmt.Println()
//					// Recreate ipStats map to reset statistics
//					ipStats = make(map[string]IPStats)
//					mutex.Unlock()
//				case <-stopCapture:
//					return
//				}
//			}
//		}(device)
//
//	}
//	// Wait for termination signal
//	<-stopCapture
//
//}

func resolveDNSName(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if err != nil {
		return "", err
	}
	// If multiple DNS names are associated with the IP address,
	// only the first one is returned.
	return names[0], nil
}

func printIpStats(stat IPStats, size string) {

	fmt.Printf("%s %-40s ", "Name:", stat.Name)
	fmt.Printf("%s %-15s ", "Host:", stat.IP)
	fmt.Printf("%s %12.2f%s ", "Sent:", formatRate(stat.Sent, size), size)
	fmt.Printf("%s %12.2f%s\n", "Received:", formatRate(stat.Received, size), size)
}

//func formatRate(bytes int, unit string) float64 {
//	bits := multiply(bytes, 8) // Convert bytes to bits
//
//	switch unit {
//	case "b":
//		return float64(bits)
//	case "B":
//		return float64(bytes)
//	case "Kb":
//		return float64(bits) / 1000 // Convert bits to kilobits
//	case "Mb":
//		return float64(bits) / 1000 / 1000 // Convert bits to megabits
//	case "Gb":
//		return float64(bits) / 1000 / 1000 / 1000 // Convert bits to gigabits
//	case "KB":
//		return float64(bytes) / 1024 // Convert bytes to kilobytes
//	case "MB":
//		return float64(bytes) / 1024 / 1024 // Convert bytes to megabytes
//	case "GB":
//		return float64(bytes) / 1024 / 1024 / 1024 // Convert bytes to gigabytes
//	default:
//		return float64(bytes)
//	}
//}

func formatRate(bytes int, unit string) float64 {
	// Define conversion factors
	const (
		bitsPerByte      = 8
		bitsPerKilobit   = 1000
		bitsPerMegabit   = bitsPerKilobit * 1000
		bitsPerGigabit   = bitsPerMegabit * 1000
		bytesPerKilobyte = 1024
		bytesPerMegabyte = bytesPerKilobyte * 1024
		bytesPerGigabyte = bytesPerMegabyte * 1024
	)

	// Convert bytes to bits
	bits := multiply(bytes, bitsPerByte)

	// Use conversion factors based on the unit
	switch unit {
	case "b":
		return float64(bits)
	case "B":
		return float64(bytes)
	case "Kb":
		return float64(bits) / bitsPerKilobit
	case "Mb":
		return float64(bits) / bitsPerMegabit
	case "Gb":
		return float64(bits) / bitsPerGigabit
	case "KB":
		return float64(bytes) / bytesPerKilobyte
	case "MB":
		return float64(bytes) / bytesPerMegabyte
	case "GB":
		return float64(bytes) / bytesPerGigabyte
	default:
		return float64(bytes)
	}
}

func multiply(x, y int) int {
	result := 0
	for i := 0; i < y; i++ {
		result += x
	}
	return result
}

func divide(dividend, divisor int) float64 {
	quotient := 0
	for dividend >= divisor {
		dividend -= divisor
		quotient++
	}

	// Calculate fractional part without using floating-point division
	if dividend > 0 {
		// Emulate fractional part using addition-based method
		remainder := multiply(dividend, 10) // Shift decimal place by multiplying by 10
		for remainder >= divisor {
			remainder -= divisor
			quotient++
		}
	}

	return float64(quotient)
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/pcap"
//	"log"
//	"net"
//	"os"
//	"os/signal"
//	"syscall"
//
//	//"os/signal"
//	"strconv"
//	"strings"
//	//"syscall"
//	"time"
//)
//
//// InterfaceStats stores the statistics for an interface
//type InterfaceStats struct {
//	Interface string  `json:"interface"`
//	Sent      float64 `json:"sent"`
//	Received  float64 `json:"received"`
//}
//
//// IPStats stores the statistics for an IP address
//type IPStats struct {
//	IP       string  `json:"ip"`
//	Sent     float64 `json:"sent"`
//	Received float64 `json:"received"`
//}
//
//func main() {
//	// Find all available network interfaces
//	devices, err := pcap.FindAllDevs()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Channel to handle termination signal
//	stopCapture := make(chan os.Signal, 1)
//	signal.Notify(stopCapture, syscall.SIGINT, syscall.SIGTERM)
//
//	// Capture packets and update IP statistics for each network interface
//	for _, device := range devices {
//		if len(device.Addresses) == 0 {
//			continue // Skip interfaces with no addresses
//		}
//
//		// Open the network interface for packet capture
//		handle, err := pcap.OpenLive(device.Name, 1600, true, pcap.BlockForever)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		defer handle.Close()
//
//		// Map to store IP address statistics
//		ipStats := make(map[string]IPStats)
//
//		// Capture packets and update IP statistics
//		go func(device pcap.Interface) {
//			for packet := range gopacket.NewPacketSource(handle, handle.LinkType()).Packets() {
//				networkLayer := packet.NetworkLayer()
//				if networkLayer == nil {
//					continue
//				}
//
//				srcIP := networkLayer.NetworkFlow().Src().String()
//				dstIP := networkLayer.NetworkFlow().Dst().String()
//
//				// Update IPStats for source IP
//				ipStats[srcIP] = IPStats{
//					IP:       srcIP,
//					Sent:     ipStats[srcIP].Sent + float64(len(packet.Data())),
//					Received: ipStats[srcIP].Received,
//				}
//
//				// Update IPStats for destination IP
//				ipStats[dstIP] = IPStats{
//					IP:       dstIP,
//					Sent:     ipStats[dstIP].Sent,
//					Received: ipStats[dstIP].Received + float64(len(packet.Data())),
//				}
//			}
//		}(device)
//
//		// Periodically print IP address statistics
//		go func(device pcap.Interface) {
//			ticker := time.NewTicker(10 * time.Second)
//			defer ticker.Stop()
//
//			for {
//				select {
//				case <-ticker.C:
//					fmt.Printf("IP Address Statistics for Interface %s:\n", device.Name)
//					for _, stats := range ipStats {
//						fmt.Printf("IP: %s | Sent: %.f bytes | Received: %.f bytes\n", stats.IP, stats.Sent, stats.Received)
//					}
//					fmt.Println()
//				case <-stopCapture:
//					return
//				}
//			}
//		}(device)
//	}
//
//	// Wait for termination signal
//	<-stopCapture
//}

//func main() {
//	// Find all available network interfaces
//	devices, err := pcap.FindAllDevs()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Capture packets and update IP statistics for each network interface
//	for _, device := range devices {
//		fmt.Printf("Device: %+v\n", device)
//		if len(device.Addresses) == 0 {
//			continue // Skip interfaces with no addresses
//		}
//
//		// Open the network interface for packet capture
//		handle, err := pcap.OpenLive(device.Name, 1600, true, pcap.BlockForever)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		defer handle.Close()
//
//		// Map to store IP address statistics
//		ipStats := make(map[string]IPStats)
//
//		// Capture packets and update IP statistics
//		go func() {
//			for packet := range gopacket.NewPacketSource(handle, handle.LinkType()).Packets() {
//				networkLayer := packet.NetworkLayer()
//				if networkLayer == nil {
//					continue
//				}
//
//				srcIP := networkLayer.NetworkFlow().Src().String()
//				dstIP := networkLayer.NetworkFlow().Dst().String()
//
//				// Update IPStats for source IP
//				ipStats[srcIP] = IPStats{
//					IP:       srcIP,
//					Sent:     float64(len(packet.Data())),
//					Received: ipStats[srcIP].Received,
//				}
//
//				// Update IPStats for destination IP
//				ipStats[dstIP] = IPStats{
//					IP:       dstIP,
//					Sent:     ipStats[dstIP].Sent,
//					Received: float64(len(packet.Data())),
//				}
//			}
//		}()
//
//		// Periodically print IP address statistics
//		ticker := time.NewTicker(10 * time.Second)
//		defer ticker.Stop()
//
//		for {
//			select {
//			case <-ticker.C:
//				fmt.Printf("IP Address Statistics for Interface %s:\n", device.Name)
//				for _, stats := range ipStats {
//					fmt.Printf("IP: %s | Sent: %.f bytes | Received: %.f bytes\n", stats.IP, stats.Sent, stats.Received)
//				}
//				fmt.Println()
//			}
//		}
//	}
//}

//func main() {
//	// Find all available network interfaces
//	devices, err := pcap.FindAllDevs()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Print information about each network interface
//	for _, device := range devices {
//		fmt.Println("Name:", device.Name)
//		fmt.Println("Description:", device.Description)
//		fmt.Println("Addresses:")
//		for _, address := range device.Addresses {
//			fmt.Printf("- IP: %s\n", address.IP)
//			fmt.Printf("  Netmask: %s\n", address.Netmask)
//		}
//		fmt.Println()
//	}
//}
//
//func main() {
//	// Define the network interface to capture packets from
//	// iface := "eth0" // Change this to the desired interface
//	iface := "en0"
//
//	// Open the network interface for packet capture
//	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer handle.Close()
//
//	// Map to store IP address statistics
//	ipStats := make(map[string]IPStats)
//
//	// Channel to handle termination signal
//	stopCapture := make(chan os.Signal, 1)
//	signal.Notify(stopCapture, syscall.SIGINT, syscall.SIGTERM)
//
//	// Capture packets and update IP statistics
//	go func() {
//		for packet := range gopacket.NewPacketSource(handle, handle.LinkType()).Packets() {
//			networkLayer := packet.NetworkLayer()
//			if networkLayer == nil {
//				continue
//			}
//
//			srcIP := networkLayer.NetworkFlow().Src().String()
//			dstIP := networkLayer.NetworkFlow().Dst().String()
//
//			// Update IPStats for source IP
//			ipStats[srcIP] = IPStats{
//				IP:       srcIP,
//				Sent:     ipStats[srcIP].Sent + float64(len(packet.Data())),
//				Received: ipStats[srcIP].Received,
//			}
//
//			// Update IPStats for destination IP
//			ipStats[dstIP] = IPStats{
//				IP:       dstIP,
//				Sent:     ipStats[dstIP].Sent,
//				Received: ipStats[dstIP].Received + float64(len(packet.Data())),
//			}
//		}
//	}()
//
//	// Periodically print IP address statistics
//	ticker := time.NewTicker(10 * time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			fmt.Println("IP Address Statistics:")
//			for _, stats := range ipStats {
//				fmt.Printf("IP: %s | Sent: %.f bytes | Received: %.f bytes\n", stats.IP, stats.Sent, stats.Received)
//			}
//			println()
//		case <-stopCapture:
//			return
//		}
//	}
//}

//func main() {
//	// Define the network interface to capture packets from
//	//iface := "eth0" // Change this to the desired interface
//	iface := "en0"
//
//	// Open the network interface for packet capture
//	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer handle.Close()
//
//	// Packet source to decode packets
//	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
//
//	// Map to store IP address statistics
//	ipStats := make(map[string]struct {
//		Sent     uint64
//		Received uint64
//	})
//
//	// Channel to handle termination signal
//	stopCapture := make(chan os.Signal, 1)
//	signal.Notify(stopCapture, syscall.SIGINT, syscall.SIGTERM)
//
//	// Capture packets and update IP statistics
//	go func() {
//
//		for packet := range packetSource.Packets() {
//			networkLayer := packet.NetworkLayer()
//			if networkLayer == nil {
//				continue
//			}
//
//			srcIP := networkLayer.NetworkFlow().Src().String()
//			dstIP := networkLayer.NetworkFlow().Dst().String()
//
//			// Update IPStats for source IP
//			if _, ok := ipStats[srcIP]; !ok {
//				ipStats[srcIP] = struct {
//					Sent     uint64
//					Received uint64
//				}{}
//			}
//			ipStats[srcIP] = struct {
//				Sent     uint64
//				Received uint64
//			}{
//				Sent:     ipStats[srcIP].Sent + uint64(len(packet.Data())),
//				Received: ipStats[srcIP].Received,
//			}
//
//			// Update IPStats for destination IP
//			if _, ok := ipStats[dstIP]; !ok {
//				ipStats[dstIP] = struct {
//					Sent     uint64
//					Received uint64
//				}{}
//			}
//			ipStats[dstIP] = struct {
//				Sent     uint64
//				Received uint64
//			}{
//				Sent:     ipStats[dstIP].Sent,
//				Received: ipStats[dstIP].Received + uint64(len(packet.Data())),
//			}
//		}
//	}()
//
//	// Periodically print IP address statistics
//	ticker := time.NewTicker(10 * time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-ticker.C:
//			fmt.Println("IP Address Statistics:")
//			for ip, stats := range ipStats {
//				fmt.Printf("IP: %s | Sent: %d bytes | Received: %d bytes\n", ip, stats.Sent, stats.Received)
//			}
//		case <-stopCapture:
//			return
//		}
//	}
//}

// readNetworkStats reads network interface statistics
// readNetworkStats reads network interface statistics

//func readNetworkStats() ([]InterfaceStats, []IPStats) {
//	file, err := os.Open("/proc/net/dev")
//	if err != nil {
//		fmt.Println("Error:", err)
//		return nil, nil
//	}
//	defer file.Close()
//
//	var interfaceStats []InterfaceStats
//	var ipStats []IPStats
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		line := scanner.Text()
//		if !strings.Contains(line, ":") {
//			continue
//		}
//
//		interfaceName, txBytes, rxBytes := parseNetworkStats(line)
//		println(interfaceName)
//		if interfaceName == "" {
//			continue
//		}
//		if interfaceName == "lo" {
//			continue
//		}
//		if txBytes == 0 && rxBytes == 0 {
//			continue
//		}
//
//		updateInterfaceStats(&interfaceStats, interfaceName, txBytes, rxBytes)
//		updateIPStats(&ipStats, interfaceName, txBytes, rxBytes)
//	}
//
//	if err := scanner.Err(); err != nil {
//		fmt.Println("Error:", err)
//	}
//
//	return interfaceStats, ipStats
//}
//
//// parseNetworkStats parses network interface statistics from a line
//func parseNetworkStats(line string) (string, int, int) {
//	fields := strings.Fields(line)
//	if len(fields) < 17 {
//		return "", 0, 0
//	}
//
//	interfaceName := strings.TrimSuffix(fields[0], ":")
//	rxBytes, _ := strconv.Atoi(fields[1])
//	txBytes, _ := strconv.Atoi(fields[9])
//
//	return interfaceName, rxBytes, txBytes
//}
//
//// updateInterfaceStats updates the interface statistics
//func updateInterfaceStats(interfaceStats *[]InterfaceStats, name string, txBytes, rxBytes int) {
//	stats := InterfaceStats{
//		Interface: name,
//		Sent:      float64(txBytes),
//		Received:  float64(rxBytes),
//	}
//	*interfaceStats = append(*interfaceStats, stats)
//}
//
//// updateIPStats updates the IP statistics
//func updateIPStats(ipStats *[]IPStats, interfaceName string, txBytes, rxBytes int) {
//	addrs, err := net.InterfaceByName(interfaceName)
//	if err != nil {
//		return
//	}
//	interfaceAddrs, _ := addrs.Addrs()
//	for _, addr := range interfaceAddrs {
//		ip := addr.(*net.IPNet).IP.String()
//		name, err := resolveDNSName(ip)
//		if err != nil {
//			println(err.Error())
//		}
//
//		stats := IPStats{
//			IP:       ip,
//			Name:     name,
//			Sent:     float64(txBytes),
//			Received: float64(rxBytes),
//		}
//		*ipStats = append(*ipStats, stats)
//	}
//}
