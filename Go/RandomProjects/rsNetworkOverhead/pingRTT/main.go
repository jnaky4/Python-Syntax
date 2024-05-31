package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func MeasureNetworkRTT(ip string, fileSizeBytes int, numPings int) (float64, error) {
	var totalRTT float64
	for i := 0; i < numPings; i++ {
		rtt, err := MeasureNetworkRTTSingle(ip, fileSizeBytes)
		if err != nil {
			return 0, err
		}
		totalRTT += rtt
	}
	averageRTT := totalRTT / float64(numPings)
	return averageRTT, nil
}

func MeasureNetworkRTTSingle(ip string, fileSizeBytes int) (float64, error) {
	cmd := exec.Command("ping", "-c", "1", "-s", strconv.Itoa(fileSizeBytes), ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}

	if !strings.Contains(string(output), "time=") {
		return 0, fmt.Errorf("no ping response")
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "time=") {
			fields := strings.Fields(line)
			for _, field := range fields {
				if strings.HasPrefix(field, "time=") {
					timeStr := strings.TrimPrefix(field, "time=")
					rtt, err := strconv.ParseFloat(timeStr, 64)
					if err != nil {
						return 0, err
					}
					fmt.Printf("Round Trip Time (RTT): %.2f ms\n", rtt)
					return rtt, nil
				}
			}
		}
	}

	fmt.Println("Error: no ping response")
	return 0, fmt.Errorf("no ping response")
}

func main() {
	ip := "8.8.8.8"       // Example IP address (Google DNS)
	fileSizeBytes := 1024 // 1 KB
	numPings := 5         // Number of times to ping the IP address

	averageRTT, err := MeasureNetworkRTT(ip, fileSizeBytes, numPings)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Average Round Trip Time (RTT) to %s with %d bytes: %.2f ms\n", ip, fileSizeBytes, averageRTT)
}
