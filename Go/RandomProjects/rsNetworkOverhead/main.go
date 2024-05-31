package main

import (
	"fmt"
)

func main() {
	// Define the size of the file in bytes
	fileSize := 1024 * 1024 // 1 MB

	// Define parameters for calculations
	distance := 100.0          // Distance in miles
	transmissionRate := 1000.0 // Transmission rate in Mbps

	// Calculate the round-trip time for each component
	physicalMediaRTT := CalculatePhysicalMediaRTT(distance)
	networkInterfaceRTT := CalculateNetworkInterfaceRTT(transmissionRate)
	tcpipOverhead := CalculateTCPIPOverhead(fileSize)
	distanceRTT := CalculateDistanceRTT(distance)

	// Calculate the total round-trip time
	totalRTT := physicalMediaRTT + networkInterfaceRTT + tcpipOverhead + distanceRTT

	// Print the total round-trip time
	fmt.Printf("Total Round-Trip Time for a %d MB file: %f ns\n", fileSize/(1024*1024), totalRTT)
}

func CalculatePhysicalMediaRTT(distanceMiles float64) float64 {
	// Speed of light in miles per hour
	speedOfLightMPH := 186282.0 // Approximately 186,282 miles per second

	// Calculate the round-trip time (RTT) based on the distance and the speed of light in mph
	// RTT = 2 * Distance / Speed_of_light * 1e9 (to convert to nanoseconds)
	rttNs := (2 * distanceMiles / speedOfLightMPH) * 3600 * 1e9 // Convert to nanoseconds

	return rttNs
}

func CalculateNetworkInterfaceRTT(transmissionRate float64) float64 {
	// Calculate round-trip time based on transmission rate
	rttNs := (float64(1) / transmissionRate) * float64(1000000000) // in ns
	return rttNs
}

func CalculateTCPIPOverhead(payloadSize int) int {
	// Ethernet frame header size
	ethernetHeaderSize := 14 // bytes

	// IPv4 header size
	ipv4HeaderSize := 20 // bytes

	// TCP header size
	tcpHeaderSize := 20 // bytes

	// Total overhead is the sum of Ethernet, IPv4, and TCP header sizes
	totalOverhead := ethernetHeaderSize + ipv4HeaderSize + tcpHeaderSize

	// Calculate the number of packets required to transmit the payload
	// Maximum payload size per packet (MTU) for Ethernet is typically 1500 bytes
	maxPayloadPerPacket := 1500 // bytes
	numPackets := payloadSize / maxPayloadPerPacket

	// Calculate the total overhead for all packets
	totalOverheadForAllPackets := totalOverhead * numPackets

	return totalOverheadForAllPackets
}

func CalculateDistanceRTT(distance float64) float64 {
	// Speed of light in miles per hour
	speedOfLightMPH := 186282 // Approximately 186,282 miles per second

	// Convert distance to miles
	distanceMiles := distance / 1.60934 // 1 mile = 1.60934 kilometers

	// Calculate the round-trip time (RTT) based on the distance and the speed of light in mph
	// RTT = 2 * Distance / Speed_of_light * 1e9 (to convert to nanoseconds)
	rttNs := (2 * distanceMiles / float64(speedOfLightMPH)) * 3600 * 1e9 // Convert to nanoseconds

	return rttNs
}
