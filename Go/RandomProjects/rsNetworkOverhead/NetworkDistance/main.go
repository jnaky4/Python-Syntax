package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

// NetworkHop represents a hop in the network topology with its IP address and latency.
type NetworkHop struct {
	IP      string
	Latency float64
}

// GetNetworkTopology finds the best route to the specified IP address and returns the hops with their latencies.
func GetNetworkTopology(ip string) ([]NetworkHop, error) {
	const maxHops = 30
	const timeoutSeconds = 5

	var bestRoute []NetworkHop
	bestLatency := float64(0)
	newBestFound := false

	fmt.Println("Finding the best route to", ip)

	// Try different hop limits until reaching the destination IP or reaching the maximum hops.
	for hopLimit := 2; hopLimit <= maxHops; hopLimit++ {
		fmt.Printf("\nTrying hop limit: %d\n", hopLimit)

		cmd := exec.Command("traceroute", "-m", strconv.Itoa(hopLimit), "-w", strconv.Itoa(timeoutSeconds), ip)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Traceroute failed:", err)
			// If traceroute fails or times out, continue to the next hop limit.
			continue
		}

		fmt.Println("\nTraceroute output:")
		fmt.Println(string(output))

		hops, err := parseTracerouteOutput(string(output))
		if err != nil {
			return nil, err
		}

		fmt.Println("\nParsed hops:")
		for _, hop := range hops {
			fmt.Printf("Hop: %s, Latency: %.2f ms\n", hop.IP, hop.Latency)
		}

		// Calculate the total latency for the route.
		totalLatency := calculateTotalLatency(hops)
		if totalLatency == 0 {
			fmt.Println("\nTotal latency is zero, destination not reached.")
			// If the total latency is zero, it indicates that the destination IP was not reached.
			continue
		}

		fmt.Printf("\nTotal latency for this route: %.2f ms\n", totalLatency)

		// Update the best route if the latency is lower than the previous best.
		if bestLatency == 0 || totalLatency < bestLatency {
			bestLatency = totalLatency
			bestRoute = hops
			newBestFound = true
			//fmt.Printf("\nNew best route found with latency: %.2f ms\n", bestLatency)
		}

		// Stop the loop if a new best route is found
		if newBestFound {
			break
		}
	}

	//fmt.Println("\nBest route to", ip, ":", bestRoute)

	return bestRoute, nil
}

// parseTracerouteOutput parses the output of the traceroute command to extract hops and latencies.
func parseTracerouteOutput(output string) ([]NetworkHop, error) {
	var hops []NetworkHop
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, " ms") {
			fields := strings.Fields(line)
			ip := fields[1]

			// Extract the latency from the last field.
			//fmt.Printf("FIELDS %+v\n", fields)

			sumLatency := 0.0
			for i := 1; i <= 3; i++ {
				latencyStr := fields[(len(fields) + (-2 * i))]
				latencyStr = strings.TrimSuffix(latencyStr, "ms") // Remove "ms" suffix
				//println("LATENCY::: ", latencyStr)
				latency, err := strconv.ParseFloat(latencyStr, 64)
				if err != nil {
					return nil, err
				}
				sumLatency += latency
			}

			hops = append(hops, NetworkHop{IP: ip, Latency: sumLatency / 3})
		}
	}
	return hops, nil
}

// calculateTotalLatency calculates the total latency for a route.
func calculateTotalLatency(hops []NetworkHop) float64 {
	var totalLatency float64
	for _, hop := range hops {
		totalLatency += hop.Latency
	}
	return totalLatency
}

// calculateDistanceFromLatency calculates the estimated distance in miles based on the total latency of the best route.
func calculateDistanceFromLatency(route []NetworkHop) float64 {
	// Speed of light in miles per millisecond (approximately)
	speedOfLightMilesPerMs := 186.282

	// Calculate the total latency for the route
	totalLatencyMs := calculateTotalLatency(route)

	// Estimate distance based on latency (RTT) and speed of light
	distance := totalLatencyMs * speedOfLightMilesPerMs

	return distance
}

func CalculateDistance() float64 {
	println()
	//ip := "161.225.130.163" // is a target ip in MN
	ip := "151.101.194.187" // Target.com ip address in SF hosted by Amazon
	//136.226.65.14 is zscalar https://whatismyipaddress.com/ip/136.226.65.14

	bestRoute, err := GetNetworkTopology(ip)
	if err != nil {
		fmt.Println("Error retrieving network topology:", err)
		return 0.0
	}

	// Calculate the estimated distance in miles based on the total latency
	estimatedDistance := calculateDistanceFromLatency(bestRoute)
	fmt.Printf("Estimated Distance to %s: %.2f miles\n", ip, estimatedDistance)
	return estimatedDistance
}

func TransferTime(fileSize float64, networkSpeedBitsPS float64) float64 {
	gbps := networkSpeedBitsPS / 1000 / 1000 / 1000

	transferTimeSeconds := CalculateTransferTime(fileSize, networkSpeedBitsPS)
	//fmt.Printf("fs %.8f bw %.8f time %.8f\n", fileSize, networkSpeedBitsPS, transferTimeSeconds)
	fmt.Printf("Transfer time on a %.2fGbps network: %.6f seconds\n", gbps, transferTimeSeconds)
	return transferTimeSeconds
}

// CalculateTransferTime calculates the transfer time given the data size in bytes and network speed in bits per second.
func CalculateTransferTime(dataSizeBytes float64, networkSpeedBitsPerSecond float64) float64 {

	transferTimeSeconds := (dataSizeBytes * 8) / networkSpeedBitsPerSecond // Convert bytes to bits and calculate transfer time
	return transferTimeSeconds
}

// CalculateBandwidth calculates the bandwidth (in Mbps) based on the data size (in bytes) and transfer time (in seconds).
func CalculateBandwidth(dataSizeBytes float64, transferTimeSeconds float64) float64 {
	// Convert dataSizeBytes to bits
	dataSizeBits := dataSizeBytes * 8

	// Calculate bandwidth in Mbps (Megabits per second)
	bandwidthMbps := dataSizeBits / (transferTimeSeconds * 1000000)

	return bandwidthMbps
}

func GetFileSize(filePath string) (float64, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Get the file information
	fileInfo, err := file.Stat()
	if err != nil {
		return 0, err
	}

	// Return the file size in bytes
	return float64(fileInfo.Size()), nil
}

func main() {

	distanceMiles := CalculateDistance()

	getwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	getwd = path.Join(getwd, "RandomProjects", "rsNetworkOverhead", "NetworkDistance")

	dataSizeNormal, err := GetFileSize(path.Join(getwd, "fullrs.json"))
	if err != nil {
		log.Fatal(err.Error())
	}
	dataSizeNormalCompressed, err := GetFileSize(path.Join(getwd, "fullrs.json.gz"))
	if err != nil {
		log.Fatal(err.Error())
	}
	dataSizeEnv, err := GetFileSize(path.Join(getwd, "env.json"))
	if err != nil {
		log.Fatal(err.Error())
	}
	dataSizeEnvCompressed, err := GetFileSize(path.Join(getwd, "env.json.gz"))
	if err != nil {
		log.Fatal(err.Error())
	}

	//dataSizeBoth := dataSizeNormalCompressed - dataSizeEnvCompressed
	dataSizeMinified := dataSizeNormal - dataSizeEnv

	numStores := 300.0
	numReplicaSets := 180.0
	waveSizeNormal := getWaveFileSize(dataSizeNormal, numReplicaSets, numStores)
	waveSizeCompressed := getWaveFileSize(dataSizeNormalCompressed, numReplicaSets, numStores)
	waveSizeMinified := getMinimizedFileSize(dataSizeMinified, dataSizeEnv, numReplicaSets, numStores)
	waveSizeBoth := getMinimizedFileSize(dataSizeNormalCompressed, dataSizeEnvCompressed, numReplicaSets, numStores)

	//fmt.Printf("Normal: %.0f bytes\nCompressed: %.0f bytes\nMinified: %.0f bytes\nBoth: %.0f bytes\n", dataSizeNormal, dataSizeNormalCompressed, dataSizeMinified, dataSizeBoth)
	fmt.Printf("Wave file size of %.0f Stores and %.0f ReplicaSets\nNormal: %.0f bytes\nCompressed: %.0f bytes\nMinified: %.0f bytes\nBoth: %.0f bytes\n", numStores, numReplicaSets, waveSizeNormal, waveSizeCompressed, waveSizeMinified, waveSizeBoth)

	//var transferTimes []float64
	BandwidthBitsPS := 1000.0 * 1000 * 1000          //1Gbps
	RealisticBandwidthBitsPS := BandwidthBitsPS * .5 // Realistic Bandwidth

	////dataSizeNormal := 1340971200.0 // Size of the data Bytes
	//transferTimes = append(transferTimes, TransferTime(dataSizeNormal, BandwidthBitsPS))
	////dataSizeMinified := 525960967.0
	//transferTimes = append(transferTimes, TransferTime(dataSizeMinified, BandwidthBitsPS))
	////dataSizeNormalCompressed := 508322188.0
	//transferTimes = append(transferTimes, TransferTime(dataSizeNormalCompressed, BandwidthBitsPS))
	////dataSizeBoth := 101664437.0
	//transferTimes = append(transferTimes, TransferTime(dataSizeBoth, BandwidthBitsPS))

	////Normal
	//transferTimes = append(transferTimes, TransferTime(waveSizeNormal, RealisticBandwidthBitsPS))
	////Minified
	//transferTimes = append(transferTimes, TransferTime(waveSizeMinified, RealisticBandwidthBitsPS))
	////Compressed
	//transferTimes = append(transferTimes, TransferTime(waveSizeCompressed, RealisticBandwidthBitsPS))
	////Both
	//transferTimes = append(transferTimes, TransferTime(waveSizeBoth, RealisticBandwidthBitsPS))

	fmt.Printf("Normal Transfer Time %.2f seconds\n", CalculateRealisticTransferTime(waveSizeNormal, RealisticBandwidthBitsPS, distanceMiles))
	fmt.Printf("Compressed Transfer Time %.2f seconds\n", CalculateRealisticTransferTime(waveSizeCompressed, RealisticBandwidthBitsPS, distanceMiles))
	fmt.Printf("Minified Transfer Time %.2f seconds\n", CalculateRealisticTransferTime(waveSizeMinified, RealisticBandwidthBitsPS, distanceMiles))
	fmt.Printf("Both Time %.2f seconds\n", CalculateRealisticTransferTime(waveSizeBoth, RealisticBandwidthBitsPS, distanceMiles))
}

func getWaveFileSize(baseFileSize float64, numReplicaSets float64, numStores float64) float64 {
	return baseFileSize * numReplicaSets * numStores
}

func getMinimizedFileSize(minimizedFileSize float64, envFileSize float64, numReplicaSets float64, numStores float64) float64 {
	return minimizedFileSize + (envFileSize * numReplicaSets * numStores)
}

// CalculateRealisticTransferTime calculates the time it takes to transfer a file over a network given the file size in bytes, bandwidth in bits per second, and distance in miles.
func CalculateRealisticTransferTime(fileSizeBytes float64, bandwidthBitsPerSecond float64, distanceMiles float64) float64 {
	// Convert bandwidth from bits per second to bytes per second
	bandwidthBytesPerSecond := bandwidthBitsPerSecond / 8

	// Calculate roundtrip time per packet
	speedOfLightMps := 299792458                                      // Speed of light in meters per second
	distanceMeters := distanceMiles * 1.60934 * 1000                  // Convert miles to meters
	roundtripTimeSeconds := distanceMeters / float64(speedOfLightMps) // Calculate roundtrip time

	// Calculate packet size including TCP/IP overhead (assuming 1500 bytes for data payload)
	packetSizeBytes := 1500.0

	// Estimate the number of packets
	numPackets := fileSizeBytes / packetSizeBytes

	// Calculate transfer time in seconds

	transferTimeSeconds := (fileSizeBytes / bandwidthBytesPerSecond) + (roundtripTimeSeconds * numPackets)

	return transferTimeSeconds
}
